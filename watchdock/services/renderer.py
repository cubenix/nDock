"""
Module to handle requests and render required template.
"""
from flask import render_template
import services.host as host
import constants


def render_dashboard(config_manager):
    hosts = config_manager.get_hosts()
    host_containers = host.get_containers(hosts)

    return render_template(
        "index.html", title="Dashboard",
        host_details=hosts,
        hosts=list(host_containers.keys()),
        containers=list(host_containers.values()),
        color_hosts=constants.COLOR_HOSTS,
        color_background=constants.COLOR_BACKGROUND,
        background_class=constants.BACKGROUND_CLASS)
