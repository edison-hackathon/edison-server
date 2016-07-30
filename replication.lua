#!/usr/bin/env tarantool

-- local fiber = require('fiber')

box.cfg{
    listen = '*:3302',
    replication_source = "100.100.147.43:3301"
}

space = box.space.measurements

require('console').start()
