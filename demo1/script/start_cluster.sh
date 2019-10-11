#!/bin/bash

$ETCD_HOME/etcd  --config-file=$ETCD_HOME/conf/conf1.yml &
$ETCD_HOME/etcd  --config-file=$ETCD_HOME/conf/conf2.yml &
$ETCD_HOME/etcd  --config-file=$ETCD_HOME/conf/conf3.yml &