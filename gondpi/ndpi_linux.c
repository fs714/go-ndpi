#include <stdio.h>
#include "ndpi_linux.h"

extern struct ndpi_flow_struct *ndpi_flow_struct_malloc()
{
    struct ndpi_flow_struct *newflow = (struct ndpi_flow_struct *)ndpi_flow_malloc(SIZEOF_FLOW_STRUCT);

    if (newflow != NULL)
    {
        memset(newflow, 0, SIZEOF_FLOW_STRUCT);
    }

    return newflow;
}

extern void ndpi_flow_struct_free(struct ndpi_flow_struct *flow)
{
    ndpi_free_flow(flow);
}

extern uint8_t ndpi_flow_get_protocol_id_already_guessed(struct ndpi_flow_struct *flow) {
    return flow->protocol_id_already_guessed;
}

extern uint8_t ndpi_flow_get_protocol_was_guessed(struct ndpi_flow_struct *flow) {
    return flow->protocol_was_guessed;
}

extern uint8_t ndpi_flow_get_init_finished(struct ndpi_flow_struct *flow) {
    return flow->init_finished;
}

extern uint8_t ndpi_flow_get_client_packet_direction(struct ndpi_flow_struct *flow) {
    return flow->client_packet_direction;
}

extern uint8_t ndpi_flow_get_packet_direction(struct ndpi_flow_struct *flow) {
    return flow->packet_direction;
}

extern uint8_t ndpi_flow_get_is_ipv6(struct ndpi_flow_struct *flow) {
    return flow->is_ipv6;
}

extern uint8_t ndpi_flow_get_first_pkt_fully_encrypted(struct ndpi_flow_struct *flow) {
    return flow->first_pkt_fully_encrypted;
}

extern uint8_t ndpi_flow_get_already_gaveup(struct ndpi_flow_struct *flow) {
    return flow->already_gaveup;
}

extern uint8_t ndpi_flow_get_extra_packets_func_set(struct ndpi_flow_struct *flow) {
    return flow->extra_packets_func != NULL;
}

extern uint32_t ndpi_flow_get_client_addr_v4(struct ndpi_flow_struct *flow) {
    return flow->c_address.v4;
}

extern uint32_t ndpi_flow_get_server_addr_v4(struct ndpi_flow_struct *flow) {
    return flow->s_address.v4;
}

extern uint8_t ndpi_proto_defaults_get_is_clear_text_proto(ndpi_proto_defaults_t *pd) {
    return pd->isClearTextProto;
}

extern uint8_t ndpi_proto_defaults_get_is_app_protocol(ndpi_proto_defaults_t *pd) {
    return pd->isAppProtocol;
}

extern uint8_t ndpi_proto_defaults_get_is_custom_proto(ndpi_proto_defaults_t *pd) {
    return pd->isCustomProto;
}

// Create a detection module, not finalized yet: configuration and custom files
// can be loaded before calling ndpi_finalize_initialization().
// Protocols whose bit is not set in detection_bitmask are disabled.
extern struct ndpi_detection_module_struct *ndpi_detection_module_create(enum ndpi_license_type license_type,
                                                                         uint32_t *detection_bitmask,
                                                                         uint32_t bitmask_len)
{
    struct ndpi_detection_module_struct *ndpi_struct = ndpi_init_detection_module(NULL, license_type);
    if (ndpi_struct == NULL)
    {
        return NULL;
    }

    char proto_id[8];
    for (uint32_t id = 1; id < bitmask_len * 32; id++)
    {
        if ((detection_bitmask[id / 32] & (1u << (id % 32))) == 0)
        {
            snprintf(proto_id, sizeof(proto_id), "%u", id);
            // Ids not matching any internal protocol are refused by nDPI: ignore them
            ndpi_set_config(ndpi_struct, proto_id, "enable", "0");
        }
    }

    return ndpi_struct;
}

extern void ndpi_detection_module_exit(struct ndpi_detection_module_struct *ndpi_struct)
{
    ndpi_exit_detection_module(ndpi_struct);
}

extern struct ndpi_proto ndpi_packet_processing(struct ndpi_detection_module_struct *ndpi_struct,
                                                struct ndpi_flow_struct *flow,
                                                const unsigned char *packet,
                                                const unsigned short packetlen,
                                                const u_int64_t packet_time_ms)
{
    struct ndpi_proto proto = ndpi_detection_process_packet(ndpi_struct, flow, packet, packetlen, packet_time_ms, NULL);

    return proto;
}
