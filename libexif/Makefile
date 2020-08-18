SOURCES=$(wildcard *.c ./canon/*.c ./fuji/*.c ./olympus/*.c ./pentax/*.c)
OBJS=$(patsubst %.c,%.o, $(SOURCES))

all:$(OBJS)
	ar -rcs libexif.a $(OBJS)

%.o:%.c
	gcc -DHAVE_CONFIG_H -I. -I.. -DGETTEXT_PACKAGE=\"libexif-12\" -Wall -Wmissing-declarations -Wmissing-prototypes -g -O2 -c $< -o $@

clean:
	rm -f $(OBJS)
	rm -f libexif.a

