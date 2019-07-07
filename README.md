These are short Go scripts that I may use to add functionality to my Go apps or for troubleshooting.

printsrv.go: probes DNS SRV and print hosts
printsrv-cli.go: same as above but it's cli version with option to add service name as parameter rather than editing file manually
getreqsrv-all.go: does same SRV request as above + issues get request to all of them (including local host)
getreq-srv-remote-only.go: optimised version of above as there is no get request for local host
