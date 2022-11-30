#include <stdio.h>
#include "ndpi_wrapper.h"

static struct ndpi_detection_module_struct *ndpi_struct = NULL;

static u_int32_t detection_tick_resolution = 1000;
static u_int32_t size_flow_struct = 0;

extern int ndpi_initialize()
{
    set_ndpi_malloc(malloc);
    set_ndpi_free(free);

    ndpi_init_prefs init_prefs = ndpi_no_prefs;
    ndpi_struct = ndpi_init_detection_module(init_prefs);
    if (ndpi_struct == NULL)
    {
        return -1;
    }

    NDPI_PROTOCOL_BITMASK protos;
    NDPI_BITMASK_SET_ALL(protos);
    ndpi_set_protocol_detection_bitmask2(ndpi_struct, &protos);

    size_flow_struct = ndpi_detection_get_sizeof_ndpi_flow_struct();

    ndpi_finalize_initialization(ndpi_struct);

    return 0;
}

extern void ndpi_exit()
{
    ndpi_exit_detection_module(ndpi_struct);
}

static struct ndpi_flow_struct *get_ndpi_flow()
{
    struct ndpi_flow_struct *newflow;
    newflow = (struct ndpi_flow_struct *)malloc(size_flow_struct);

    if (newflow == NULL)
    {
        printf("[NDPI] %s(1): not enough memory\n", __FUNCTION__);
        return NULL;
    }

    memset(newflow, 0, size_flow_struct);

    return newflow;
}

static void free_ndpi_flow(struct ndpi_flow_struct *flow)
{
    ndpi_free(flow);
}

static int packet_processing(struct ndpi_flow_struct *ndpi_flow, const struct ndpi_iphdr *iph, u_int16_t ipsize,
                             const u_int64_t time)
{
    u_int16_t protocol = 0;

    ndpi_protocol detected = ndpi_detection_process_packet(ndpi_struct, ndpi_flow, (uint8_t *)iph, ipsize, time);
    protocol = detected.master_protocol;
    if (protocol == 0)
    {
        protocol = detected.app_protocol;
    }

    return protocol;
}

int main()
{
}
