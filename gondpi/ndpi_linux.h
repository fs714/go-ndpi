#include <ndpi/ndpi_main.h>

extern struct ndpi_flow_struct *ndpi_flow_struct_malloc();
extern void ndpi_flow_struct_free(struct ndpi_flow_struct *);
extern uint8_t ndpi_flow_get_protocol_id_already_guessed(struct ndpi_flow_struct *);
extern uint8_t ndpi_flow_get_protocol_was_guessed(struct ndpi_flow_struct *);
extern uint8_t ndpi_flow_get_init_finished(struct ndpi_flow_struct *);
extern uint8_t ndpi_flow_get_client_packet_direction(struct ndpi_flow_struct *);
extern uint8_t ndpi_flow_get_packet_direction(struct ndpi_flow_struct *);
extern uint8_t ndpi_flow_get_is_ipv6(struct ndpi_flow_struct *);
extern uint8_t ndpi_flow_get_first_pkt_fully_encrypted(struct ndpi_flow_struct *);
extern uint8_t ndpi_flow_get_already_gaveup(struct ndpi_flow_struct *);
extern uint8_t ndpi_flow_get_extra_packets_func_set(struct ndpi_flow_struct *);
extern uint32_t ndpi_flow_get_client_addr_v4(struct ndpi_flow_struct *);
extern uint32_t ndpi_flow_get_server_addr_v4(struct ndpi_flow_struct *);

extern uint8_t ndpi_proto_defaults_get_is_clear_text_proto(ndpi_proto_defaults_t *);
extern uint8_t ndpi_proto_defaults_get_is_app_protocol(ndpi_proto_defaults_t *);
extern uint8_t ndpi_proto_defaults_get_is_custom_proto(ndpi_proto_defaults_t *);

extern struct ndpi_detection_module_struct *ndpi_detection_module_create(enum ndpi_license_type, uint32_t *, uint32_t);
extern void ndpi_detection_module_exit(struct ndpi_detection_module_struct *);
extern struct ndpi_proto ndpi_packet_processing(struct ndpi_detection_module_struct *,
                                                struct ndpi_flow_struct *,
                                                const unsigned char *,
                                                const unsigned short,
                                                const u_int64_t);
