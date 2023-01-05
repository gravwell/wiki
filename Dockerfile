FROM python:3.10-buster

RUN apt update -y
RUN apt install -y enchant

RUN pip3 install sphinx==5.1.1 myst-parser==0.18.0 pydata-sphinx-theme==0.11.0 sphinx-design==0.3.0 sphinxcontrib-spelling==7.6.2
