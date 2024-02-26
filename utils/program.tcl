# If no command line arguments are passed perform no actions
if {$argc > 0} {
# argv0 conatins the filename of the script
puts "The name of the script is: $argv0"
# argc contains the count of the arguments passed
puts "Total count of arguments passed is: $argc"
# argv contains a list of the arguments
puts "The arguments passed are: $argv"
# Using the List Index of argv print a specific argument
puts "The first argument passed was [lindex $argv 0]"
} else {
    puts "NO ARGS PASSED"
}

# Extract the firmwarename from the command-line arguments
set firmwarename [lindex $argv 0]

# Print the firmwarename
puts $firmwarename

open_hw_manager
connect_hw_server
close_hw_target
open_hw_target [lindex [get_hw_targets] 0]
set_property PROBES.FILE {} [lindex [get_hw_devices] 0]

# Print the firmwarename
puts "Firmware Name: $firmwarename"
set_property PROGRAM.FILE "/app/bitstreams/__BOARDNUMBER__/$firmwarename" [lindex [get_hw_devices] 0]

program_hw_devices [lindex [get_hw_devices] 0]
refresh_hw_device [lindex [get_hw_devices] 0]
exit