import os

from flask import Flask, request, jsonify
from sqlalchemy import create_engine, func
from sqlalchemy.orm import sessionmaker

from recommendations.model import Base, Rating

DATABASE_URL = os.environ.get('DATABASE_URL', 'sqlite:///ratings.db')
engine = create_engine(DATABASE_URL)
Session = sessionmaker(bind=engine)
Base.metadata.create_all(engine)  # Migrate

app = Flask(__name__)


@app.route('/api/ratings/', methods=['POST'])
def add_rating():
    """
    Create a new rating for a meal.  Request body should be json in the below format:
    (yaml representation, square brackets means optional)
    user_id: int
    menu_id: int
    restaurant_id: int
    ratings:
        [portion_size]: 1 <= int <= 10
        [healthiness]: 1 <= int <= 10
        [sweetness]: 1 <= int <= 10
        [spice_level]: 1 <= int <= 10
        [rating]: 1 <= int <= 10
    """
    data = request.get_json()
    REQUIRED_PARAMS = frozenset(['user_id', 'menu_id', 'restaurant_id', 'ratings'])
    missing_params = [param for param in REQUIRED_PARAMS if param not in data]
    if missing_params:
        raise ValueError('Missing required parameters from request body: {}'.format(missing_params))
    session = Session()
    rating = Rating.create(data['user_id'], data['menu_id'], data['restaurant_id'], **data['ratings'])
    session.add(rating)
    session.commit()
    return jsonify(rating.serialize())


@app.route('/api/ratings/', methods=['GET'])
def get_rating():
    """
    Get the average ratings for a given menu_id
    :query_param menu: the id of the meal in question
    :returns List of average ratings objects (list of length 1 when menu param is set) for each menu
    """
    menu_id = request.args.get('menu')
    restaurant_id = request.args.get('restaurant')
    if menu_id is None and restaurant_id is None:
        raise ValueError('A menu or restaurant must be provided in the query string e.g. /api/ratings/?menu=21')
    session = Session()
    avg_ratings = session.query(
        Rating.menu_id,
        func.count(Rating.menu_id).label('count'),
        func.avg(Rating.portion_size).label('portion_size'),
        func.avg(Rating.healthiness).label('healthiness'),
        func.avg(Rating.sweetness).label('sweetness'),
        func.avg(Rating.spice_level).label('spice_level'),
        func.avg(Rating.rating).label('rating'),
    )
    if menu_id is not None:
        avg_ratings = avg_ratings.filter(Rating.menu_id == menu_id)
    else:
        avg_ratings = avg_ratings.filter(Rating.restaurant_id == restaurant_id).group_by(Rating.menu_id)
    return jsonify([_serialize_ratings_row(row) for row in avg_ratings])


def _serialize_ratings_row(row):
    """
    Converts the namedtuple-like row that sqlalchemy returns into a dict that resembles the object used to create
    a rating.
    """
    return {
        'menu_id': row.menu_id,
        'count': row.count,
        'ratings': {key: getattr(row, key) for key in row._fields if key not in ('menu_id', 'count')}
    }


# http://flask.pocoo.org/docs/0.12/server/#in-code
if __name__ == '__main__':
    app.run()
