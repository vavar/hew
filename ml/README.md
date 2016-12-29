## Recommendations Microservice

This is a microservice for making meal recommendations.  It uses Machine Learning to improve recommendations over time and is written in Python.

### Install

- First ensure that you have a working python 3 virtualenv.  You will also need a Fortran compiler (!), for that simply run `brew install gcc` on OSX.
- Once you have activated your virtualenv, run `pip install -r requirements.txt` to install the dependencies.  On my machine I had to install scipy first with `pip install scipy`.

### Run
- To run the service, run `python recommentations.py`.  It will bind to port 8080 on 127.0.0.1 (only accessible from localhost).
- When the service starts, it will open a sqlite db in the same directory and rebuild its models from the latest data.
- Once the models are up-to-date, you will be able to query the service for recommendations.
