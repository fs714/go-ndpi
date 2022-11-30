#include <ndpi/ndpi_main.h>

extern int ndpi_initialize();
extern void ndpi_exit();
static struct ndpi_flow_struct *get_ndpi_flow();
static void free_ndpi_flow(struct ndpi_flow_struct*);
static int packet_processing(struct ndpi_flow_struct*, const struct ndpi_iphdr*, u_int16_t, const u_int64_t);
