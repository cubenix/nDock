from flask import Flask
from config_manager import ConfigManager
import services.renderer as renderer


app = Flask(__name__)
hosts = ConfigManager().get_hosts()


@app.route('/')
@app.route('/index')
@app.route('/home')
@app.route('/index.html')
def index():
    return renderer.render_dashboard(hosts)


@app.route('/stats')
def stats():
    return renderer.render_stats(hosts[0])


@app.route('/get_stats')
def get_stats():
    return renderer.get_stats(hosts[0])


if __name__ == '__main__':
    app.run(debug=True)
