require 'digest'
require 'erb'
require 'git'
require 'octokit'
require 'open-uri'
require 'ostruct'
require 'rake'
require 'rake_circle_ci'
require 'rake_github'
require 'rake_gpg'
require 'rake_ssh'
require 'semantic'
require 'yaml'

require_relative 'lib/platform'
require_relative 'lib/version'

task :default => %w(clean all:format all:vet cli:build:all test:end_to_end:run)

task :clean do
  puts "Cleaning temporary directories..."
  rm_rf('build/*')
  Dir.glob("work/*")
      .select { |file| /^[^.]/.match file }
      .each { |file| rm_rf(file) }
  puts
end

namespace :encryption do
  namespace :directory do
    desc 'Ensure CI secrets directory exists'
    task :ensure do
      FileUtils.mkdir_p('secrets/ci')
    end
  end

  namespace :passphrase do
    desc 'Generate encryption passphrase for CI GPG key'
    task generate: ['directory:ensure'] do
      File.open('secrets/ci/encryption.passphrase', 'w') do |f|
        f.write(SecureRandom.base64(36))
      end
    end
  end
end

namespace :keys do
  namespace :deploy do
    namespace :gorgon do
      RakeSSH.define_key_tasks(
        path: 'secrets/ci/',
        name_prefix: 'gorgon.ssh',
        comment: 'tobyclemson@gmail.com'
      )
    end

    namespace :homebrew_utils do
      RakeSSH.define_key_tasks(
        path: 'secrets/ci/',
        name_prefix: 'homebrew-utils.ssh',
        comment: 'tobyclemson@gmail.com'
      )
    end
  end

  namespace :secrets do
    namespace :gpg do
      RakeGPG.define_generate_key_task(
        output_directory: 'secrets/ci',
        name_prefix: 'gpg',
        owner_name: 'Toby Clemson',
        owner_email: 'tobyclemson@gmail.com',
        owner_comment: 'gorgon CI Key'
      )
    end

    task generate: ['gpg:generate']
  end
end

namespace :secrets do
  desc 'Regenerate all secrets'
  task regenerate: %w[
    encryption:passphrase:generate
    keys:deploy:gorgon:generate
    keys:deploy:homebrew_utils:generate
    keys:secrets:generate
  ]
end

RakeCircleCI.define_project_tasks(
  namespace: :circle_ci,
  project_slug: 'github/tobyclemson/gorgon'
) do |t|
  circle_ci_config =
    YAML.load_file('secrets/circle_ci/config.yaml')

  t.api_token = circle_ci_config['circle_ci_api_token']
  t.environment_variables = {
    ENCRYPTION_PASSPHRASE:
      File.read('secrets/ci/encryption.passphrase')
          .chomp
  }
  t.checkout_keys = []
  t.ssh_keys = [
    {
      hostname: 'github.com',
      private_key: File.read('secrets/ci/gorgon.ssh.private')
    },
    {
      hostname: 'github.com',
      private_key: File.read('secrets/ci/homebrew-utils.ssh.private')
    }
  ]
end

namespace :github do
  RakeGithub.define_repository_tasks(
    namespace: :gorgon,
    repository: 'tobyclemson/gorgon',
    ) do |t|
    github_config =
      YAML.load_file('secrets/github/credentials.yaml')

    t.access_token = github_config['github_token']
    t.deploy_keys = [
      {
        title: 'CircleCI',
        public_key: File.read('secrets/ci/gorgon.ssh.public')
      }
    ]
  end

  RakeGithub.define_repository_tasks(
    namespace: :homebrew_utils,
    repository: 'tobyclemson/homebrew-utils',
    ) do |t|
    github_config =
      YAML.load_file('secrets/github/credentials.yaml')

    t.access_token = github_config['github_token']
    t.deploy_keys = [
      {
        title: 'CircleCI - Gorgon',
        public_key: File.read('secrets/ci/homebrew-utils.ssh.public')
      }
    ]
  end
end

namespace :pipeline do
  desc 'Prepare CircleCI Pipeline'
  task prepare: %i[
    circle_ci:project:follow
    circle_ci:env_vars:ensure
    circle_ci:checkout_keys:ensure
    circle_ci:ssh_keys:ensure
    github:gorgon:deploy_keys:ensure
    github:homebrew_utils:deploy_keys:ensure
  ]
end

namespace :tools do
  namespace :install do
    desc "Install gox"
    task :gox do
      puts "Installing gox..."
      sh('bash -c "go get github.com/mitchellh/gox"')
      puts
    end

    task :all => %w(tools:install:gox)
  end
end

namespace :dependencies do
  desc "Vendor all dependencies"
  task :vendor do
    puts "Vendoring dependencies..."
    sh('bash -c "go mod vendor"')
    puts
  end
end

namespace :cli do
  desc "Vet the CLI tool source"
  task :vet => %w(dependencies:vendor) do
    packages = go_packages_satisfying(
        exclusions: %w(vendor test))
    puts "Vetting production code..."
    sh("bash -c \"go vet #{packages}\"")
    puts
  end

  desc "Format the CLI tool source"
  task :format do
    packages = go_packages_satisfying(
        exclusions: %w(vendor test))
    puts 'Formatting production code...'
    sh("bash -c \"go fmt #{packages}\"")
    puts
  end

  namespace :build do
    desc "Build the CLi tool for OS X"
    task :darwin,
        [:version] => %w(tools:install:all dependencies:vendor) do |_, args|
      version = args.version || next_prerelease_version

      build_version(version, os_arches: 'darwin/amd64')
    end

    desc "Build the CLI tool for all OSs and architectures"
    task :all,
        [:version] => %w(tools:install:all dependencies:vendor) do |_, args|
      version = args.version || next_prerelease_version

      build_version(version)
    end
  end

  desc "Pre-release a new version of the CLI tool"
  task :prerelease do
    version = next_prerelease_version

    Rake::Task["cli:build:all"].invoke(version)

    release_version(version, prerelease: true)
  end

  desc "Release a new version of the CLI tool"
  task :release do
    version = next_release_version

    Rake::Task["cli:build:all"].invoke(version)

    release_version(version, prerelease: false)
  end
end

namespace :test do
  namespace :end_to_end do
    desc 'Vet the end-to-end test source'
    task :vet => %w(dependencies:vendor) do
      packages = go_packages_satisfying(
          inclusions: %w(test))
      puts "Vetting end-to-end test code..."
      sh("bash -c \"go vet #{packages}\"")
      puts
    end

    desc 'Format the end-to-end test source'
    task :format do
      packages = go_packages_satisfying(
          inclusions: %w(test))
      puts 'Formatting end-to-end test code...'
      sh("bash -c \"go fmt #{packages}\"")
      puts
    end

    task :run, [:version] do |_, args|
      packages = go_packages_satisfying(
          inclusions: %w(test/end_to_end))

      github_credentials = YAML.load_file('secrets/github/credentials.yaml')
      github_token = github_credentials['github_token']
      ENV["TEST_GITHUB_TOKEN"] ||= github_token

      ssh_private_key_path = "#{Dir.getwd}/secrets/ssh/gorgon.ssh.private"
      ENV["TEST_SSH_PRIVATE_KEY_PATH"] ||= ssh_private_key_path

      binary_os = Platform.os
      binary_architecture = Platform.architecture
      binary_version = args.version || next_prerelease_version
      binary_directory = "#{binary_version}_#{binary_os}_#{binary_architecture}"
      binary_path = "build/bin/#{binary_directory}/gorgon"
      binary_path_env_var = "TEST_BINARY_PATH=#{binary_path}"

      puts "Running end-to-end tests..."
      sh("bash -c \"#{binary_path_env_var} go test -v #{packages}\"")
      puts
    end
  end
end

namespace :homebrew do
  namespace :formula do
    desc "Generate homebrew formula for a new version of the CLI tool"
    task :generate, [:version] do |_, args|
      version = args.version || latest_version

      puts "Generating homebrew formula for version: #{version}"
      url = "https://github.com/tobyclemson/gorgon/archive/#{version}.tar.gz"

      mkdir_p("build/dist")
      open("build/dist/#{version}.tar.gz", 'wb') do |file|
        file << open(url).read
      end

      checksum = Digest::SHA256.file("build/dist/#{version}.tar.gz")

      template = ERB.new(File.read("templates/formula.rb.erb"))

      mkdir_p("build/formula")
      open('build/formula/gorgon.rb', 'w') do |file|
        file << template.result(
            OpenStruct
                .new(version: version, checksum: checksum)
                .instance_eval { binding })
      end
    end

    desc "Push a new homebrew formula for a new version of the CLI tool"
    task :push, [:version] do |_, args|
      version = args.version || next_release_version

      Rake::Task['homebrew:formula:generate'].invoke(version)

      puts "Pushing homebrew formula for version: #{version}"
      mkdir_p "build/repos"
      repo = Git.clone(
          'git@github.com:tobyclemson/homebrew-utils.git',
          'homebrew-utils',
          path: 'build/repos')

      FileUtils.cp('build/formula/gorgon.rb', 'build/repos/homebrew-utils')

      repo.config('user.name', 'CI')
      repo.config('user.email', 'ci@tobyclemson.com')

      repo.commit("Updating formula for version #{version}", all: true)
      repo.push
    end
  end
end

namespace :all do
  desc "Vet all source"
  task :vet => %w(cli:vet test:end_to_end:vet)

  desc "Format all source"
  task :format => %w(cli:format test:end_to_end:format)
end

def repo
  Git.open('.')
end

def latest_version
  repo.tags.map do |tag|
    Semantic::Version.new(tag.name)
  end.max
end

def next_prerelease_version
  latest_version.rc!.to_s
end

def next_release_version
  latest_version.release!.to_s
end

def build_version(version, options = {})
  os_arches = options[:os_arches] ||
      'linux/amd64 darwin/amd64 windows/amd64'
  output_dir = "build/bin/#{version}_{{.OS}}_{{.Arch}}/{{.Dir}}"
  package = "github.com/tobyclemson/gorgon"

  osarch_switch = "-osarch='#{os_arches}'"
  output_switch = "-output='#{output_dir}'"
  ldflags_switch = "-ldflags '-X main.Version=#{version}'"
  switches = "#{osarch_switch} #{output_switch} #{ldflags_switch}"

  puts "Building CLI with version: #{version}..."
  sh("bash -c \"gox #{switches} #{package}\"")
  puts
end

def release_version(version, options = {})
  github_credentials = YAML.load_file('secrets/github/credentials.yaml')
  github_token = github_credentials['github_token']

  binary_specifier = ->(v, os, arch) {
    return "#{v}_#{os}_#{arch}"
  }
  binary_path = ->(v, os, arch, ext) {
    ext_part = ext ? ".#{ext}" : ""
    return "build/bin/#{binary_specifier.call(v, os, arch)}/gorgon#{ext_part}"
  }

  puts "Releasing CLI with version: #{version}"
  client = Octokit::Client.new(access_token: github_token)
  release = client.create_release('tobyclemson/gorgon', version,
      name: version,
      draft: true,
      prerelease: options[:prerelease] || false)
  [
      [:darwin, :amd64, nil],
      [:linux, :amd64, nil],
      [:windows, :amd64, :exe]
  ].each do |os_arch|
    client.upload_asset(
        release.url,
        binary_path.call(version, os_arch[0], os_arch[1], os_arch[2]),
        name: "gorgon-#{version}-#{os_arch[0]}-#{os_arch[1]}",
        content_type: "application/octet-stream")
  end
  client.update_release(release.url, draft: false)

  puts
end

def go_packages_satisfying(exclusions: [], inclusions: [])
  grep_exclusions = exclusions.collect { |f| "grep -v #{f}" }.join(' | ')
  grep_inclusions = inclusions.collect { |f| "grep #{f}" }.join(' | ')

  grep_invocations = [grep_exclusions, grep_inclusions].
      flatten.
      delete_if(&:empty?).
      join(' | ')

  command = "go list ./... | #{grep_invocations}"

  `bash -c "#{command}"`.gsub("\n", ' ')
end