#include "ndpi_linux.h"

extern struct ndpi_detection_module_struct *ndpi_struct_initialize()
{
    set_ndpi_malloc(malloc);
    set_ndpi_free(free);

    struct ndpi_detection_module_struct *ndpi_struct;
    ndpi_init_prefs init_prefs = ndpi_no_prefs;
    ndpi_struct = ndpi_init_detection_module(init_prefs);
    if (ndpi_struct == NULL)
    {
        return NULL;
    }

    NDPI_PROTOCOL_BITMASK protos;
    NDPI_BITMASK_SET_ALL(protos);
    ndpi_set_protocol_detection_bitmask2(ndpi_struct, &protos);

    ndpi_finalize_initialization(ndpi_struct);

    return ndpi_struct;
}

extern void ndpi_struct_exit(struct ndpi_detection_module_struct *ndpi_struct)
{
    ndpi_exit_detection_module(ndpi_struct);
}

extern struct ndpi_flow_struct *get_ndpi_flow()
{
    struct ndpi_flow_struct *newflow;
    newflow = (struct ndpi_flow_struct *)ndpi_flow_malloc(SIZEOF_FLOW_STRUCT);

    if (newflow != NULL)
    {
        memset(newflow, 0, SIZEOF_FLOW_STRUCT);
    }

    return newflow;
}

extern void free_ndpi_flow(struct ndpi_flow_struct *flow)
{
    ndpi_free_flow(flow);
}

extern struct ndpi_proto ndpi_packet_processing(struct ndpi_detection_module_struct *ndpi_struct,
                                                struct ndpi_flow_struct *flow,
                                                const unsigned char *packet,
                                                const unsigned short packetlen,
                                                const u_int64_t packet_time_ms)
{
    // printf("------\n");
    // printf("%p\n", ndpi_struct);
    // printf("%p\n", flow);
    // printf("%p\n", packet);
    // printf("%d\n", packetlen);
    // printf("%ld\n", packet_time_ms);

    struct ndpi_proto proto;
    proto = ndpi_detection_process_packet(ndpi_struct, flow, packet, packetlen, packet_time_ms);

    return proto;
}

extern u_int16_t ndpi_category_to_id(ndpi_protocol_category_t category) {
    u_int16_t id = (u_int16_t)category;
    return id;
}
