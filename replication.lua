#!/usr/bin/env tarantool

-- local fiber = require('fiber')

box.cfg{
    listen = '*:3302',
    replication_source = "100.100.147.43:3301"
}

devices = box.space.devices

require('console').start()
