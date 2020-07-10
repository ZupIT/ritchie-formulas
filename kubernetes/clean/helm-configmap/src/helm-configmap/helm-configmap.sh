#!/bin/sh
run() {
  CONFIGMAP_LIST=$(kubectl -n kube-system get configmap |awk '{print $1}'|cut -f1 -d.|sort|uniq|grep -v 'NAME');
  for i in $CONFIGMAP_LIST;
  do echo "$i";
  for j in $(kubectl -n kube-system get configmap | grep "$i" | awk '{print $1}'|cut -d. -f2|cut -dv -f2| sort -n | head -n -2); do echo "$i".v"$j"; kubectl -n kube-system delete configmap "$i".v"$j" ; done; done
}
