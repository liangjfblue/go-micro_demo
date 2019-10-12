#!/bin/bash

#3个节点http://172.16.7.16:9002 http://172.16.7.16:9004 http://172.16.7.16:9006

#存放etcd集群数据
mkdir -p $ETCD_HOME/data/data1
mkdir -p $ETCD_HOME/data/data2
mkdir -p $ETCD_HOME/data/data3

#存放etcd集群配置文件
mkdir -p $ETCD_HOME/conf
touch $ETCD_HOME/conf/conf1.yml
touch $ETCD_HOME/conf/conf2.yml
touch $ETCD_HOME/conf/conf3.yml

echo "name: node1
data-dir: /home/liangjf/app/etcd-v3.3.13-linux-amd64/data/data1
listen-client-urls: 'http://172.16.7.16:9002'
advertise-client-urls: 'http://172.16.7.16:9002'
listen-peer-urls: 'http://172.16.7.16:9001'
initial-advertise-peer-urls: 'http://172.16.7.16:9001'
initial-cluster: node1=http://172.16.7.16:9001,node2=http://172.16.7.16:9003,node3=http://172.16.7.16:9005
initial-cluster-token: etcd-cluster-1
initial-cluster-state: new" > $ETCD_HOME/conf/conf1.yml

echo "name: node2
data-dir: /home/liangjf/app/etcd-v3.3.13-linux-amd64/data/data2
listen-client-urls: 'http://172.16.7.16:9004'
advertise-client-urls: 'http://172.16.7.16:9004'
listen-peer-urls: 'http://172.16.7.16:9003'
initial-advertise-peer-urls: 'http://172.16.7.16:9003'
initial-cluster: node1=http://172.16.7.16:9001,node2=http://172.16.7.16:9003,node3=http://172.16.7.16:9005
initial-cluster-token: etcd-cluster-1
initial-cluster-state: new" > $ETCD_HOME/conf/conf2.yml

echo "name: node3
data-dir: /home/liangjf/app/etcd-v3.3.13-linux-amd64/data/data3
listen-client-urls: 'http://172.16.7.16:9006'
advertise-client-urls: 'http://172.16.7.16:9006'
listen-peer-urls: 'http://172.16.7.16:9005'
initial-advertise-peer-urls: 'http://172.16.7.16:9005'
initial-cluster: node1=http://172.16.7.16:9001,node2=http://172.16.7.16:9003,node3=http://172.16.7.16:9005
initial-cluster-token: etcd-cluster-1
initial-cluster-state: new" > $ETCD_HOME/conf/conf3.yml