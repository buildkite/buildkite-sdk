require_relative "lib/buildkite/version"

Gem::Specification.new do |spec|
  spec.name = "buildkite-sdk"
  spec.version = Buildkite::VERSION
  spec.authors = ["Buildkite"]
  spec.email = ["support@buildkite.com"]

  spec.summary = "A Ruby SDK for Buildkite!"
  spec.description = "A Ruby SDK for Buildkite."
  spec.homepage = "https://buildkite.com"
  spec.required_ruby_version = ">= 3.0.0"

  spec.metadata["homepage_uri"] = spec.homepage
  spec.metadata["source_code_uri"] = "https://github.com/buildkite/buildkite-sdk"
  spec.metadata["changelog_uri"] = "https://github.com/buildkite/buildkite-sdk"

  gemspec = File.basename(__FILE__)

  spec.files = IO.popen(%w[git ls-files -z], chdir: __dir__, err: IO::NULL) do |ls|
    ls.readlines("\x0", chomp: true).reject do |f|
      (f == gemspec) ||
        f.start_with?(*%w[bin/ test/ spec/ features/ .git appveyor Gemfile])
    end
  end

  spec.bindir = "exe"
  spec.executables = spec.files.grep(%r{\Aexe/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]
  spec.add_dependency "ostruct", "~> 0.1.0"
end
