"""
Module to handle requests and render required template.
"""
from flask import render_template
import services.host as host
import services.container as container
import constants


def render_dashboard(hosts):
    host_containers = host.get_containers(hosts)

    return render_template(
        "index.html", title="Dashboard",
        host_details=hosts,
        hosts=list(host_containers.keys()),
        containers=list(host_containers.values()),
        color_hosts=constants.COLOR_HOSTS,
        color_background=constants.COLOR_BACKGROUND,
        background_class=constants.BACKGROUND_CLASS)


def render_stats(docker_host, hosts):
    containers = host.get_containers_and_color_codes(docker_host)
    return render_template(
        "stats.html",
        containers=containers,
        docker_host=docker_host,
        host_details=hosts)


def get_stats(host, container_id):
    return container.get_stats(host, container_id)
