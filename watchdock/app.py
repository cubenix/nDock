from quart import Quart
# import asyncio
from config_manager import ConfigManager
import services.renderer as renderer


app = Quart(__name__)
hosts = ConfigManager().get_hosts()


@app.route('/')
@app.route('/index')
@app.route('/home')
@app.route('/index.html')
def index():
    return renderer.render_dashboard(hosts)


@app.route('/stats')
def stats():
    return renderer.render_stats(hosts)


@app.route('/get_stats')
def get_stats():
    return renderer.get_stats(hosts[:1])


if __name__ == '__main__':
    app.run(debug=True)
