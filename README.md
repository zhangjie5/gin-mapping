# gin-mapping
A tool that register router in gin easily

### Usage

this tool will mapping handle path in this way:

- use Group() method return value as group
- use method name as handle path

finally will get `/group/method_name` in gin routers

usage can refer `mapping_test.go` file.
