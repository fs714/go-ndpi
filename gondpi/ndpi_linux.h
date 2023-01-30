#include <ndpi/ndpi_main.h>

extern void ndpi_protocol_bitmask_add(NDPI_PROTOCOL_BITMASK *, uint16_t);
extern void ndpi_protocol_bitmask_del(NDPI_PROTOCOL_BITMASK *, uint16_t);
extern bool ndpi_protocol_bitmask_is_set(NDPI_PROTOCOL_BITMASK *, uint16_t);
extern void ndpi_protocol_bitmask_reset(NDPI_PROTOCOL_BITMASK *);
extern void ndpi_protocol_bitmask_set_all(NDPI_PROTOCOL_BITMASK *);

extern struct ndpi_detection_module_struct *ndpi_detection_module_initialize(ndpi_init_prefs, NDPI_PROTOCOL_BITMASK *);
extern void ndpi_detection_module_exit(struct ndpi_detection_module_struct *);
extern ndpi_proto_defaults_t *ndpi_proto_defaults_get(struct ndpi_detection_module_struct *, bool *, bool *);

extern struct ndpi_flow_struct *ndpi_flow_struct_malloc();
extern void ndpi_flow_struct_free(struct ndpi_flow_struct *);

extern struct ndpi_proto ndpi_packet_processing(struct ndpi_detection_module_struct *,
                                                struct ndpi_flow_struct *,
                                                const unsigned char *,
                                                const unsigned short,
                                                const u_int64_t);
