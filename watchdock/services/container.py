import services.api as api


def get_stats(host, container_id):
    client = api.create_low_level_client(host)
    stats = client.stats(container_id, stream=False)
    client.close()
    return "{0:.3f}".format(__calculate_usage(stats))


def __calculate_usage(stats):
    # get required readings
    cpu_stats = stats["cpu_stats"]["cpu_usage"]["total_usage"]
    precpu_stats = stats["precpu_stats"]["cpu_usage"]["total_usage"]
    system_cpu_stats = stats["cpu_stats"]["system_cpu_usage"]
    system_precpu_stats = stats["precpu_stats"]["system_cpu_usage"]
    cpu_count = len(stats["precpu_stats"]["cpu_usage"]["percpu_usage"])

    # calculate the change for cpu usage of the container in between readings
    cpu_delta = cpu_stats - precpu_stats

    # calculate the change for the entire system between readings
    system_delta = system_cpu_stats - system_precpu_stats

    cpu_usage = 0.0
    if system_delta > 0 and cpu_delta > 0:
        cpu_usage = (cpu_delta / system_delta) * cpu_count * 100
    return cpu_usage
