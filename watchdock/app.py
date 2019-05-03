from flask import Flask, request
import json
from config_manager import ConfigManager
import services.renderer as renderer
from models.host import Host


app = Flask(__name__)
config_manager = ConfigManager()
hosts = config_manager.get_hosts()
host_dict = config_manager.get_host_dict()


@app.route('/')
@app.route('/index')
@app.route('/home')
@app.route('/index.html')
def index():
    return renderer.render_dashboard(hosts)


@app.route('/stats/<hostname>')
def stats(hostname):
    host = Host(name=hostname, IP=host_dict[hostname])
    return renderer.render_stats(host, hosts)


@app.route('/get_stats')
def get_stats():
    hostname = request.args.get('hostname')
    container_id = request.args.get('containerId')
    host = Host(hostname, host_dict[hostname])
    return json.dumps({
        'usage': renderer.get_stats(host, container_id),
        'cIndex': request.args.get('cIndex')
    })


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
