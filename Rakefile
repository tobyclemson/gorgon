require 'rake'

task :default => %w(clean cli:format cli:vet cli:build)

task :clean do
  rm_rf('build/*')
  Dir.glob("work/*")
      .select{ |file| /^[^.]/.match file }
      .each { |file| rm_rf(file)}
end

namespace :tools do
  namespace :install do
    desc "Install gox"
    task :gox do
      puts "Installing gox..."
      sh('bash -c "go get github.com/mitchellh/gox"')
    end

    task :all => %w(tools:install:gox)
  end
end

namespace :dependencies do
  desc "Vendor all dependencies"
  task :vendor do
    puts "Vendoring dependencies..."
    sh('bash -c "go mod vendor"')
  end
end

namespace :cli do
  desc "Vet the CLI tool source"
  task :vet => %w(dependencies:vendor) do
    packages = go_packages_satisfying(
        exclusions: %w(vendor test))
    puts "Vetting production code..."
    sh("bash -c \"go vet #{packages}\"")
  end

  desc "Format the CLI tool source"
  task :format do
    packages = go_packages_satisfying(
        exclusions: %w(vendor test))
    puts 'Formatting production code...'
    sh("bash -c \"go fmt #{packages}\"")
  end

  desc "Build the CLI tool"
  task :build => %w(tools:install:all dependencies:vendor) do
    version = ENV['VERSION'] || 'local'
    os_arches = 'linux/amd64 darwin/amd64'
    output_dir = "build/bin/#{version}_{{.OS}}_{{.Arch}}/{{.Dir}}"
    package = "github.com/tobyclemson/gorgon"

    osarch_switch = "-osarch='#{os_arches}'"
    output_switch = "-output='#{output_dir}'"

    puts "Building CLI with version: #{version}..."
    sh("bash -c \"gox #{osarch_switch} #{output_switch} #{package}\"")
  end
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