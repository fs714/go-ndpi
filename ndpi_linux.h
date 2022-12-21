#include <ndpi/ndpi_main.h>

extern struct ndpi_detection_module_struct *ndpi_struct_initialize();
extern void ndpi_struct_exit(struct ndpi_detection_module_struct *);
extern struct ndpi_flow_struct *get_ndpi_flow();
extern void free_ndpi_flow(struct ndpi_flow_struct *);
extern struct ndpi_proto ndpi_packet_processing(struct ndpi_detection_module_struct *,
                                                struct ndpi_flow_struct *,
                                                const unsigned char *,
                                                const unsigned short,
                                                const u_int64_t);
