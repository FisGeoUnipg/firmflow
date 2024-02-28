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