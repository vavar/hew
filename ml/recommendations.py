import os
from typing import Mapping

from flask import Flask, request, jsonify
from sqlalchemy import create_engine, Column, Integer
from sqlalchemy.sql import func
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

DATABASE_URL = os.environ.get('DATABASE_URL', 'sqlite:///ratings.db')
Base = declarative_base()
Session = sessionmaker()


class Rating(Base):
    """
    A single user's rating of a single meal.
    """
    __tablename__ = 'rating'

    id = Column(Integer, primary_key=True)
    user_id = Column(Integer, index=True, nullable=False)
    menu_id = Column(Integer, index=True, nullable=False)
    restaurant_id = Column(Integer, index=True, nullable=False)
    portion_size = Column(Integer)
    healthiness = Column(Integer)
    sweetness = Column(Integer)
    spice_level = Column(Integer)
    rating = Column(Integer)

    def __str__(self):
        return f'Rating of meal {self.menu_id} by user {self.user_id}: {self.rating}'

    def serialize(self):
        return {
            'id': self.id,
            'user_id': self.user_id,
            'menu_id': self.menu_id,
            'restaurant_id': self.restaurant_id,
            'ratings': {
                'portion_size': self.portion_size,
                'healthiness': self.healthiness,
                'sweetness': self.sweetness,
                'spice_level': self.spice_level,
                'rating': self.rating,
            },
        }


def rate_menu(user_id: int, menu_id: int, restaurant_id: int, **ratings: Mapping[str, int]):
    """
    Add a rating for the attributes of this meal
    :param user_id: Id of the submitting user
    :param menu_id: Id of the meal
    :param restaurant_id: Id of the restaurant
    :param ratings: Map of ratings. Possible keys are:
        portion_size
        healthiness
        sweetness
        spice_level
        rating
    All keys are optional, and any provided should be integer scores 1-10
    """
    if any(1 > int(score) > 10 for score in ratings.values()):
        raise ValueError('Scores should be integers from 1-10 inclusive.')

    ALLOWED_ATTRS = frozenset(['portion_size', 'healthiness', 'sweetness', 'spice_level', 'rating'])
    invalid_attrs = [attr for attr in ratings.keys() if attr not in ALLOWED_ATTRS]
    if invalid_attrs:
        raise ValueError('Invalid attributes submitted: {}'.format(invalid_attrs))

    return Rating(user_id=user_id, menu_id=menu_id, restaurant_id=restaurant_id, **ratings)


engine = create_engine(DATABASE_URL)
Base.metadata.create_all(engine)  # Migrate
Session.configure(bind=engine)
app = Flask(__name__)


@app.route('/api/ratings/', methods=['POST'])
def add_rating():
    data = request.get_json()
    REQUIRED_PARAMS = frozenset(['user_id', 'menu_id', 'restaurant_id', 'ratings'])
    missing_params = [param for param in REQUIRED_PARAMS if param not in data]
    if missing_params:
        raise ValueError('Missing required parameters from request body: {}'.format(missing_params))
    session = Session()
    rating = rate_menu(data['user_id'], data['menu_id'], data['restaurant_id'], **data['ratings'])
    session.add(rating)
    session.commit()
    return jsonify(rating.serialize())


@app.route('/api/ratings/', methods=['GET'])
def get_rating():
    """
    Get the average ratings for a given menu_id
    :query_param menu: the id of the meal in question
    """
    menu_id = request.args.get('menu')
    if menu_id is None:
        raise ValueError('A menu id must be provided in the query string e.g. /api/ratings/?menu=21')
    session = Session()
    avg_ratings = session.query(
        func.avg(Rating.portion_size).label('portion_size'),
        func.avg(Rating.healthiness).label('healthiness'),
        func.avg(Rating.sweetness).label('sweetness'),
        func.avg(Rating.spice_level).label('spice_level'),
        func.avg(Rating.rating).label('rating'),
    ).filter(Rating.menu_id == menu_id).first()
    return jsonify({key: getattr(avg_ratings, key) for key in avg_ratings._fields})


# http://flask.pocoo.org/docs/0.12/server/#in-code
if __name__ == '__main__':
    app.run()
