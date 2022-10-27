clean-primary:
	rm -rf dbs-primary .litefs-primary
	mkdir -p dbs-primary

clean-replica:
	rm -rf dbs-replica .litefs-replica
	mkdir -p dbs-replica

primary: clean-primary
	litefs -config litefs-primary.yml

replica: clean-replica
	litefs -config litefs-replica.yml
