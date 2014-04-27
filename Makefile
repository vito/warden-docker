all:
	mkdir -p skeleton/bin
	cd src && make clean all
	cp src/wsh/wshd skeleton/bin
	cp src/wsh/wsh skeleton/bin
	cp src/oom/oom skeleton/bin
	cp src/iomux/iomux-spawn skeleton/bin
	cp src/iomux/iomux-link skeleton/bin
	cp src/repquota/repquota bin
