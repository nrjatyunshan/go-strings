#include <gelf.h>
#include <libelf.h>
#include <unistd.h>
#include <stdio.h>
#include <assert.h>
#include <fcntl.h>
#include <string.h>

static void split_strings(const char *buffer, const size_t len)
{
	size_t curlen = 1;
	const char *pos = buffer;
	while (pos + 2 * curlen < buffer + len && curlen < 50)
	{
		printf("%.*s\n", curlen, pos);
		if (strncmp(pos, pos + curlen, curlen) < 0)
		{
			pos += curlen;
			continue;
		}
		if (memchr(pos, 0x0a, 2 * curlen))
		{
			pos += curlen;
			continue;
		}

		pos += curlen;
		++curlen;
	}
}

int main()
{
	int fd;
	Elf *e;
	Elf_Data *d;

	assert(elf_version(EV_CURRENT) != EV_NONE);

	fd = open("../hello/hello", O_RDONLY, 0);
	assert(fd >= 0);

	e = elf_begin(fd, ELF_C_READ, NULL);
	assert(e);

	size_t begin = 0x973f8;
	size_t end = 0x9f820;
	d = elf_getdata_rawchunk(e, begin, end - begin, ELF_T_BYTE);
	assert(d);

	split_strings(d->d_buf, d->d_size);

	elf_end(e);
	close(fd);
}
