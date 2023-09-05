FROM summerwind/actions-runner:v2.308.0-ubuntu-22.04

USER root
RUN curl -sSL https://install.python-poetry.org | python3 - 

USER runner
RUN poetry --version