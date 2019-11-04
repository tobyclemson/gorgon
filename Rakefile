require 'rake'
require 'yaml'
require 'git'
require 'semantic'
require 'octokit'

require_relative 'lib/platform'
require_relative 'lib/version'

task :default => %w(clean all:format all:vet cli:build test:end_to_end:run)

task :clean do
  puts "Cleaning temporary directories..."
  rm_rf('build/*')
  Dir.glob("work/*")
      .select { |file| /^[^.]/.match file }
      .each { |file| rm_rf(file) }
  puts
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

  desc "Build the CLI tool"
  task :build,
      [:version] => %w(tools:install:all dependencies:vendor) do |_, args|
    version = args.version || next_prerelease_version
    os_arches = 'linux/amd64 darwin/amd64 windows/amd64'
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

  desc "Release a new version of the CLI tool"
  task :release do
    version = next_release_version

    github_credentials = YAML.load_file('secrets/github/credentials.yaml')
    github_token = github_credentials['github_token']

    binary_specifier = ->(v, os, arch) {
      return "#{v}_#{os}_#{arch}"
    }
    binary_path = ->(v, os, arch, ext) {
      ext_part = ext ? ".#{ext}" : ""
      return "build/bin/#{binary_specifier.call(v, os, arch)}/gorgon#{ext_part}"
    }

    Rake::Task["cli:build"].invoke(version)

    puts "Releasing CLI with version: #{version}"
    client = Octokit::Client.new(access_token: github_token)
    release = client.create_release('tobyclemson/gorgon', version,
        name: version,
        draft: true)
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
      github_token_env_var = "TEST_GITHUB_TOKEN=#{github_token}"

      binary_os = Platform.os
      binary_architecture = Platform.architecture
      binary_version = args.version || next_prerelease_version
      binary_directory = "#{binary_version}_#{binary_os}_#{binary_architecture}"
      binary_path = "build/bin/#{binary_directory}/gorgon"
      binary_path_env_var = "TEST_BINARY_PATH=#{binary_path}"

      env_vars = "#{github_token_env_var} #{binary_path_env_var}"

      puts "Running end-to-end tests..."
      sh("bash -c \"#{env_vars} go test -v #{packages}\"")
      puts
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