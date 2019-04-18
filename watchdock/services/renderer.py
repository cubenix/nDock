"""
Module to handle requests and render required template.
"""
from flask import render_template
import services.host as host
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


def render_stats(host):
    stats = [
        {
            "name": "container-one",
            "usage": [12, 23, 14, 35, 27, 50, 65, 85, 123, 63.6, 156, 78]
        },
        {
            "name": "container-two",
            "usage": [5, 35, 10, 25, 47, 36, 41, 65, 51, 73, 91, 132]
        }
    ]
    return render_template("stats.html", stats=stats)


# def render_stats(host):
#     return render_template(
#         "stats.html",
#         color_bkg=constants.COLOR_BACKGROUND,
#         containers = containers)
# hostname=hostname)

def get_stats(hosts):
    return host.get_stats(hosts)
