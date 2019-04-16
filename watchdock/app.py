from flask import Flask, render_template
from config_manager import ConfigManager
import services.host as host
import constants


app = Flask(__name__)
config_manager = ConfigManager()


@app.route('/')
@app.route('/index')
@app.route('/home')
@app.route('/index.html')
def index():
    hosts = config_manager.get_hosts()
    host_containers = host.get_containers(hosts)

    return render_template(
        "index.html", title="Dashboard",
        hosts=list(host_containers.keys()),
        containers=list(host_containers.values()),
        color_hosts=constants.COLOR_HOSTS,
        color_background=constants.COLOR_BACKGROUND)


if __name__ == '__main__':
    app.run(debug=True)
