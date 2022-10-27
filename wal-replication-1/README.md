## The Error

Eventually it hits an error like this:

    2022-10-27 23:20:25.141257453 +0000 UTC m=+12.359139025 3498
    2022-10-27 23:20:25.143373436 +0000 UTC m=+12.361255010 3498
    2022/10/27 23:20:25 database disk image is malformed
    exit status 1

## To Reproduce

run the primary and the replica:

	  make primary
	  make replica

then run the producer and consumer:

	  go run . producer
	  go run . consumer
