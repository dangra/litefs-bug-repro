clean-primary:
	rm -rf dbs-primary .litefs-primary
	mkdir -p dbs-primary

clean-replica:
	rm -rf dbs-replica .litefs-replica
	mkdir -p dbs-replica

litefs-primary: clean-primary
	litefs -config litefs-primary.yml

litefs-replica: clean-replica
	litefs -config litefs-replica.yml
