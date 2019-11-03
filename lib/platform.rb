require 'rubygems'

module Platform
  class <<self
    def os
      case Gem::Platform.local.os
      when "darwin"
        "darwin"
      when "linux"
        "linux"
      when "mingw32"
        "windows"
      else
        raise "Unknown platform OS: #{Gem::Platform.local.os}"
      end
    end

    def architecture
      case Gem::Platform.local.cpu
      when "x86_64"
        "amd64"
      else
        raise "Unknown platform architecture: #{Gem::Platform.local.cpu}"
      end
    end
  end
end