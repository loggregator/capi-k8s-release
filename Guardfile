# A sample Guardfile
# More info at https://github.com/guard/guard#readme

## Uncomment and set this to only include directories you want to watch
directories %w(src/cloud_controller_ng) \
  .select{|d| Dir.exist?(d) ? d : UI.warning("Directory #{d} does not exist")}

## Note: if you are using the `directories` clause above and you are not
## watching the project directory ('.'), then you will want to move
## the Guardfile to a watched dir and symlink it back, e.g.
#
#  $ mkdir config
#  $ mv Guardfile config/
#  $ ln -s config/Guardfile .
#
# and, you'll have to watch "config/Guardfile" instead of "Guardfile"
guard :shell do
  watch(/.*/) do |m|
    puts "🔍 Changes to #{m[0]} detected..."
    puts "😱 Live Reloading!"

    if system('./dev/build-and-rollout-capi.sh')
      "😎 Live Reload succeeded."
    else
      "😡 Live Reload FAILED!"
    end
  end
end
