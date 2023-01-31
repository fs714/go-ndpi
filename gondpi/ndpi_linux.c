#include "ndpi_linux.h"

extern void ndpi_protocol_bitmask_add(NDPI_PROTOCOL_BITMASK *bitmask, uint16_t proto)
{
    NDPI_BITMASK_ADD(*bitmask, proto);
}

extern void ndpi_protocol_bitmask_del(NDPI_PROTOCOL_BITMASK *bitmask, uint16_t proto)
{
    NDPI_BITMASK_DEL(*bitmask, proto);
}

extern bool ndpi_protocol_bitmask_is_set(NDPI_PROTOCOL_BITMASK *bitmask, uint16_t proto)
{
    return NDPI_ISSET(bitmask, proto);
}

extern void ndpi_protocol_bitmask_reset(NDPI_PROTOCOL_BITMASK *bitmask)
{
    NDPI_BITMASK_RESET(*bitmask);
}

extern void ndpi_protocol_bitmask_set_all(NDPI_PROTOCOL_BITMASK *bitmask)
{
    NDPI_BITMASK_SET_ALL(*bitmask);
}

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

extern struct ndpi_detection_module_struct *ndpi_detection_module_initialize(ndpi_init_prefs prefs, NDPI_PROTOCOL_BITMASK *detection_bitmask)
{
    set_ndpi_malloc(malloc);
    set_ndpi_free(free);

    struct ndpi_detection_module_struct *ndpi_struct = ndpi_init_detection_module(prefs);
    if (ndpi_struct == NULL)
    {
        return NULL;
    }

    // printf("bitmask length = %ld\n", sizeof(detection_bitmask->fds_bits) / sizeof(detection_bitmask->fds_bits[0]));
    // for (int i = 0; i < sizeof(detection_bitmask->fds_bits) / sizeof(detection_bitmask->fds_bits[0]); i++)
    // {
    //     printf("bitmask %d: %u\n", i, detection_bitmask->fds_bits[i]);
    // }

    ndpi_set_protocol_detection_bitmask2(ndpi_struct, detection_bitmask);

    ndpi_finalize_initialization(ndpi_struct);

    return ndpi_struct;
}

extern void ndpi_detection_module_exit(struct ndpi_detection_module_struct *ndpi_struct)
{
    ndpi_exit_detection_module(ndpi_struct);
}

extern ndpi_proto_defaults_t *ndpi_proto_defaults_get(struct ndpi_detection_module_struct *ndpi_struct,
                                                      bool *is_clear_text_proto, bool *is_app_protocol)
{
    ndpi_proto_defaults_t *pd = ndpi_get_proto_defaults(ndpi_struct);

    for (uint32_t i = 0; i < NDPI_MAX_SUPPORTED_PROTOCOLS + NDPI_MAX_NUM_CUSTOM_PROTOCOLS; i++)
    {
        is_clear_text_proto[i] = pd[i].isClearTextProto;
        is_app_protocol[i] = pd[i].isAppProtocol;
    }

    return pd;
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

    struct ndpi_proto proto = ndpi_detection_process_packet(ndpi_struct, flow, packet, packetlen, packet_time_ms);

    return proto;
}
