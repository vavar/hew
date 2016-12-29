import os
from typing import Mapping

from sqlalchemy import create_engine, Column, Integer
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


def rate_menu(user_id: int, menu_id: int, attrs: Mapping[str, int]):
    """
    Add a rating for the attributes of this meal
    :param user_id: Id of the submitting user
    :param menu_id: Id of the meal
    :param attrs: Map of attributes. Possible attributes are:
        portion_size
        healthiness
        sweetness
        spice_level
        rating
    All attributes are optional, and all should be integer scores 1-10
    """
    if any(1 < int(score) > 10 for score in attrs.values()):
        raise ValueError('Scores should be integers from 1-10 inclusive.')

    ALLOWED_ATTRS = frozenset(['portion_size', 'healthiness', 'sweetness', 'spice_level', 'rating'])
    invalid_attrs = [attr for attr in attrs.keys() if attr not in ALLOWED_ATTRS]
    if invalid_attrs:
        raise ValueError('Invalid attributes submitted: {}'.format(invalid_attrs))

    rating = Rating(user_id, menu_id, **attrs)


if __name__ == '__main__':
    engine = create_engine(DATABASE_URL)
    Base.metadata.create_all(engine)  # Migrate
    Session.configure(bind=engine)
