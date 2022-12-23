clean-primary:
	rm -rf dbs-primary .litefs-primary
	mkdir -p dbs-primary

clean-replica:
	rm -rf dbs-replica .litefs-replica
	mkdir -p dbs-replica

primary: clean-primary
	-umount -l dbs-primary
	litefs mount -config litefs-primary.yml

replica: clean-replica
	-umount -l dbs-primary
	litefs mount -config litefs-replica.yml
