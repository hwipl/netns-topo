# netns-topo

netns-topo is a command line tool for deploying network topologies on a local
host in linux, e.g., for testing purposes. Topologies are defined as yaml
files. Nodes are deployed as network namespaces. Links between nodes are
deployed as veth interfaces. Special node types are "bridge" and "router".

## Installation

You can download, build and install netns-topo with its dependencies to your
GOPATH or GOBIN (usually `~/go/bin/`) with the go tool:

```console
$ go install github.com/hwipl/netns-topo/cmd/netns-topo
```

## Usage

You can run netns-topo with the following command:

```console
$ netns-topo
```

Make sure your user has permissions to create network namespaces, veth
interfaces and to change the network configuration on the local host.

Command line arguments of `netns-topo`:

```
Usage of netns-topo:
  netns-topo <command>

Commands:
  start <topology> [force]
        start topology
  stop <topology> [force]
        stop topology
  list
        list topologies
  run <topology> <node> <cmd>
        run command on node in topology
  save <topology>
        save topology
  remove <topology>
        remove saved topology
  help
        show this help

<topology> is a yaml file or the name of a currently active or saved topology.
<node> is the name of the node as defined in the topology.
<cmd> is a regular linux command like "ip route".
```

For example topology files, see examples below.

## Examples

Deploying a topology defined in a file called `mytopo.yaml`:

```console
$ sudo netns-topo start mytopo.yaml
```

Tearing down the topology with the name `MyTopo`:

```console
$ sudo netns-topo stop MyTopo
```

Running the command `ip route` on node `Node1` in topology `MyTopo`:

```console
$ sudo netns-topo run MyTopo Node1 "ip route"
```

Network topology with a bridge:

```yaml
---
name: TopoBridge1
nodes:
    - name: Node1
      type: node
    - name: Node2
      type: bridge
    - name: Node3
      type: node
links:
    - name: Link1
      type: veth
      nodes:
          - Node1
          - Node2
      ips:
          - 192.168.1.1/24
          - ""
    - name: Link2
      type: veth
      nodes:
          - Node2
          - Node3
      ips:
          - ""
          - 192.168.1.2/24
run:
    - node: Node1
      commands:
          - ping -c 3 192.168.1.2
```

Network topology with a router:

```yaml
---
name: TopoRouter1
nodes:
    - name: Node1
      type: node
      routes:
          - route: 192.168.2.0/24
            via: 192.168.1.1
      run:
          - ip route
    - name: Node2
      type: router
    - name: Node3
      type: node
      routes:
          - route: 192.168.1.0/24
            via: 192.168.2.1
      run:
          - ip route
links:
    - name: Link1
      type: veth
      nodes:
          - Node1
          - Node2
      macs:
          - 1e:2e:3e:4e:5e:01
          - 1e:2e:3e:4e:5e:02
      ips:
          - 192.168.1.2/24
          - 192.168.1.1/24
    - name: Link2
      type: veth
      nodes:
          - Node2
          - Node3
      ips:
          - 192.168.2.1/24
          - 192.168.2.2/24
run:
    - node: Node1
      commands:
          - ping -c 3 192.168.2.2
```
