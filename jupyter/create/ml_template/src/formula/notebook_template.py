from enum import Enum


class Estimator(Enum):
    SVC = "SVC"
    KNeighborsClassifier = "KNeighborsClassifier"
    RandomForestClassifier = "RandomForestClassifier"
    SGDClassifier = "SGDClassifier"
    SGDRegressor = "SGDRegressor"
    SVR = "SVR"
    MiniBatchKMeans = "MiniBatchKMeans"
    KMeans = "KMeans"


content = {}
content[Estimator.SVC.value] = {
    "import": "from sklearn.svm import SVC",
    "doc_link": "https://scikit-learn.org/stable\
    /modules/generated/sklearn.svm.SVC.html#sklearn.svm.SVC",
    "evaluation": "classification",
}
content[Estimator.KNeighborsClassifier.value] = {
    "import": "from sklearn.neighbors import KNeighborsClassifier",
    "doc_link": "https://scikit-learn.org/stable/modules/neighbors.html",
    "evaluation": "classification",
}
content[Estimator.RandomForestClassifier.value] = {
    "import": "from sklearn.ensemble import RandomForestClassifier",
    "doc_link": "https://scikit-learn.org/stable\
    /modules/generated/sklearn.ensemble.Random\
    ForestClassifier.html#sklearn.ensemble.RandomForestClassifier",
    "evaluation": "classification",
}
content[Estimator.SGDClassifier.value] = {
    "import": "from sklearn.linear_model import SGDClassifier",
    "doc_link": "https://scikit-learn.org/\
    stable/modules/generated/sklearn.linear\
    _model.SGDClassifier.html#sklearn.\
    linear_model.SGDClassifier",
    "evaluation": "classification",
}
content[Estimator.SGDRegressor.value] = {
    "import": "from sklearn.linear_model import SGDRegressor",
    "doc_link": "https://scikit-learn.org\
    /stable/modules/generated/sklearn.linear\
    _model.SGDRegressor.html#sklearn.linear_model.SGDRegressor",
    "evaluation": "regression",
}
content[Estimator.SVR.value] = {
    "import": "from sklearn.svm import SVR",
    "doc_link": "https://scikit-learn.org/\
    stable/modules/generated/sklearn.svm.\
    SVR.html#sklearn.svm.SVR",
    "evaluation": "regression",
}
content[Estimator.MiniBatchKMeans.value] = {
    "import": "from sklearn.cluster import MiniBatchKMeans",
    "doc_link": "https://scikit-learn.org/\
    stable/modules/clustering.html#mini-batch-k-means",
    "evaluation": "classification",
}
content[Estimator.KMeans.value] = {
    "import": "from sklearn.cluster import KMeans",
    "doc_link": "https://scikit-learn.org/\
    stable/modules/clustering.html#k-means",
    "evaluation": "clustering",
}


metrics = {
    "regression": [
        {
            "cell_type": "markdown",
            "metadata": {},
            "source": [
                "## Regression Evaluation Metrics\n",
                "\n",
                "\n",
                "Here are three common evaluation \
                metrics for regression problems:\n",
                "\n",
                "**Mean Absolute Error** (MAE) is the mean\
                 of the absolute value of the errors:\n",
                "\n",
                "$$\\frac 1n\\sum_{i=1}^n|y_i-\\hat{y}_i|$$\n",
                "\n",
                "**Mean Squared Error** (MSE) is the mean \
                of the squared errors:\n",
                "\n",
                "$$\\frac 1n\\sum_{i=1}^n(y_i-\\hat{y}_i)^2$$\n",
                "\n",
                "**Root Mean Squared Error** (RMSE) is the \
                square root of the mean of the squared errors:\n",
                "\n",
                "$$\\sqrt{\\frac 1n\\sum_{i=1}^n(y_i-\\hat{y}_i)^2}$$\n",
                "\n",
                "Comparing these metrics:\n",
                "\n",
                "- **MAE** is the easiest to understand,\
                 because it's the average error.\n",
                '- **MSE** is more popular than MAE, because \
                MSE "punishes" larger errors, which tends to \
                be useful in the real world.\n',
                '- **RMSE** is even more popular than MSE, \
                because RMSE is interpretable in the "y" units.\n',
                "\n",
                "All of these are **loss functions**, \
                because we want to minimize them.",
            ],
        },
        {
            "cell_type": "code",
            "execution_count": 0,
            "metadata": {},
            "outputs": [],
            "source": [
                "from sklearn import metrics\n",
                "print('MAE:', metrics.mean_absol\
                ute_error(y_test, predictions))\n",
                "print('MSE:', metrics.mean_squared\
                _error(y_test, predictions))\n",
                "print('RMSE:', np.sqrt(metrics.mean\
                _squared_error(y_test, predictions)))",
            ],
        },
    ],
    "classification": [
        {
            "cell_type": "markdown",
            "metadata": {},
            "source": [
                "## Classification Evaluation Metrics\n",
                "\n",
                "\n",
                "For classification problems, we might\
                 have a more binary decision. \
                Either we predict correctly or not. \
                Let's suppose we are predicting whether a certain profile of \
                person will be able to pay a loan back or not. \
                We might actually encounter some cases:\n",
                "\n",
                "* **True positive**: we have correctly predicted that\
                 the person will be able to pay the loan back\n",
                "* **False positive**: we have *incorrectly* predicted \
                the person will be able to pay the loan back\n",
                "* **True negative**: we have correctly predicted that\
                 the person will *not* be able to pay the loan back\n",
                "* **False negative**: we have *incorrectly* predicted\
                 the person will *not* be able to pay the loan back\n",
                "\n",
                "These cases may lead us to build a **confusion matrix**,\
                 which helps us understand the \
                 accuracy of our predictions. \n",
                "\n",
                '<img width=700 src="https://prod\
                -images-static.radiopaedia\
                .org/images/49024440/0f59a975b60e83f5309a5f59075e7f\
                _jumbo.jpeg" />\n',
                "\n",
                "From these values, we can infer the following metrics:\n",
                "\n",
                "**Accuracy** How many predictions were \
                right compared to the total\n",
                "\n",
                "$$\\frac {TP + TN}{Total}$$\n",
                "\n",
                "**Precision** Ability to find relevant cases in dataset\n",
                "\n",
                "$$\\frac {TP}{TP + FP}$$\n",
                "\n",
                "**Recall** Ability to only find relevant data points\n",
                "\n",
                "$$\\frac {TP}{TP + FN}$$\n",
                "\n",
                "**F1** Harmonic mean of recall and precision. \
                The harmonic mean punishes extreme values better\n",
                "\n",
                "$$2 * \\frac {Precision * Recall}{Precision + Recall}$$\n",
                "\n",
                "Great, with these terms in mind, \
                we can evaluate our classification models",
            ],
        },
        {
            "cell_type": "code",
            "execution_count": 0,
            "metadata": {},
            "outputs": [],
            "source": [
                "from sklearn.metrics import \
                classification_report, confusion_matrix\n",
                "predictions = pipe.predict(x_test)\n",
                "\n",
                "print(classification_report(predictions, y_test))\n",
                "print(confusion_matrix(predictions, y_test))",
            ],
        },
    ],
    "clustering": [
        {
            "cell_type": "markdown",
            "metadata": {},
            "source": [
                "## Clustering metrics\n",
                "\n",
                "Unfortunately, there is no clear way to validate clustering \
                data since the learning is unsupervised.\
                 You could run <a href='https://www.\
                 geeksforgeeks.org/elbow-method-for-optimal\
                 -value-of-k-in-kmeans/'> \
                 an elbow method</a> to find an optimal value or \
                 use field-related knowledge to decide on the best\
                  number of clusters. Use the `n_clusters` \
                  parameter in the model to adjust",
            ],
        }
    ],
}


def grid_search(estimator):
    if estimator not in [Estimator.SVC.value, Estimator.SVR.value]:
        return []

    return [
        {
            "cell_type": "markdown",
            "metadata": {},
            "source": [
                "# Gridsearch\n",
                "\n",
                "Finding the right parameters (like what C or gamma values to use) \
                is a tricky task. We can adopt a trial and error approach to \
                find the best fit. Through GridSearch, we can try different \
                combinations of parameters and roll with the best option. \
                You just need to feed a dictionary with possible parameters \
                and Scikit-learn will use the one with \
                the best score on the next train fit!",
            ],
        },
        {
            "cell_type": "code",
            "execution_count": 0,
            "metadata": {},
            "outputs": [],
            "source": [
                "from sklearn.model_selection import GridSearchCV\n",
                "param_grid = dict(C=[0.1,1,10,1000], \
                gamma=[1,0.1,0.01,0.001,0.0001], \
                kernel=['rbf', 'linear'])\n",
                "grid = GridSearchCV({}(), param_grid,\
                 refit=True, verbose=3)\n".format(
                    estimator
                ),
                "grid.fit(x_train, y_train)\n",
                "print('Best estimator: ' + grid.best_estimator_)\n",
                "prediction2 = grid.predict(x_test)\n",
                "print(classification_report(pred2, y_test))\n",
                "print(confusion_matrix(pred2, y_test))\n",
            ],
        },
    ]


def render_template(estimator):

    return {
        "cells": [
            {
                "cell_type": "markdown",
                "metadata": {},
                "source": [
                    "___\n",
                    "\n",
                    "<a href='https://ritchiecli.io/'> \
                    <img src='https://i.ibb.co/MBHbCfv/ritchie-logo.png'\
                     width=200 /></a>\n",
                    "___\n",
                    "# {} with Ritchie\n".format(estimator),
                    "\n",
                    "Welcome to Ritchie's generated machine learning notebook!\
                     <a href='https://scikit-learn.org/stable/tutorial\
                     /machine_learning_map/index.html'>\
                     From your choices,</a> Ritchie was able to structure an\
                      initial setup for your data analysis, but we are far\
                       from finished! Keep in mind some of the \
                       responsibilities on your hands:",
                    "\n",
                    "1. You need to define your datasource.\
                     Whether you are loading a file or connecting to\
                      a remote database, we have \
                      provided you with some common methods",
                    "\n",
                    "2. Feature engineer your data! Most of the times\
                     you will have to shape your data accordingly \
                     to feed your model. The dedicated section\
                      has more information on it ;)",
                    "\n",
                    "3. Play with your estimators' parameters! \
                    On some cases we will try to help out, \
                    such as setting up Grid Searches, but many \
                    times you will have to play with the values \
                    yourself to find the best fit. Also, you do not \
                    need to take the given estimator as final, \
                    <a href='https://scikit-learn.org/stable/\
                    tutorial/machine_learning_map/index.html'>\
                    you can check the cheatsheet</a> on other\
                     possible candidates as well.",
                    "\n",
                    "4. Test and validate your data! We provide some \
                    ways to collect metrics to validade your data, but \
                    feel free to add more to it!",
                    "\n",
                    "5. Use your model! After training it, \
                    export and go use on some real world problems!",
                    "\n\n"
                    "<a href={}> You can find the documentation \
                    for this estimator here </a>".format(
                        content[estimator]["doc_link"]
                    ),
                ],
            },
            # Data import
            {
                "cell_type": "markdown",
                "metadata": {},
                "source": [
                    "## Import data\n",
                    "\n",
                    "Import here your data in any format. \
                    You might `read_cvs`, `read_xls`, \
                    fetch from a remote database etc.",
                ],
            },
            {
                "cell_type": "code",
                "execution_count": 0,
                "metadata": {},
                "outputs": [],
                "source": [
                    "import pandas as pd\n",
                    "\n",
                    "# --- if you have a csv ---\n",
                    "# df = pd.read_csv('<YOUR FILE PATH>')\n",
                    "\n",
                    "\n",
                    "# --- if you have a xls ---\n",
                    "# df = pd.read_excel('<YOUR FILE PATH>')\n",
                    "\n",
                    "\n",
                    "\n",
                    "# --- if you are connecting to a mongo database ---\n",
                    "# from pymongo import MongoClient\n",
                    "\n",
                    "# def read_mongo(db, \
                    collection, query={}, \
                    host='localhost', port=27017, \
                    username=None, password=None):\n",
                    '#     """ Read from Mongo and \
                    Store into DataFrame """\n',
                    "\n",
                    "#     # Connect to MongoDB\n",
                    "#     mongo_uri = 'mongodb://None:None@\
                    localhost:27017/db'\n",
                    "#     db = MongoClient(mongo_uri)[db]\n",
                    "\n",
                    "#     # Make a query to the specific DB and Collection\n",
                    "#     cursor = db[collection].find(query)\n",
                    "\n",
                    "#     # Expand the cursor and construct the DataFrame\n",
                    "#     df =  pd.DataFrame(list(cursor))\n",
                    "\n",
                    "#     return df\n",
                    "\n",
                    "\n",
                    "# --- if you are connecting to a postgres database ---\n",
                    "# import psycopg2\n",
                    "\n",
                    '# conn = psycopg2.connect(host="localhost",\
                     port = 5432, database="suppliers", \
                     user="postgres", password="postgres")\n',
                    "\n",
                    "# # Create a cursor object\n",
                    "# cur = conn.cursor()\n",
                    "\n",
                    '# # A sample query of all data from the "vendors"\
                     table in the "suppliers" database\n',
                    '# cur.execute("""SELECT * FROM vendors""")\n',
                    "# query_results = cur.fetchall()",
                ],
            },
            # Feature engineering
            {
                "cell_type": "markdown",
                "metadata": {},
                "source": [
                    "## Feature engineering\n",
                    "\n",
                    "This is the part in which you can do feature engineering. \
                    Feature engineering is the ability to extract and format \
                    more information from the data given field knowledge. \
                    For instance, you might want to extract the month of \
                    the year to calculate house pricing given you know that\
                     year seasonality might affect it. Some manipulations you \
                     can do with your data are:\n",
                    "\n",
                    "* Remove non-numerical columns\n",
                    '* <a href="https://pandas.pydata.org/pandas\
                    -docs/stable/reference/api/pandas.get_dummies.html">\
                    Dummify your categorical data</a>\n',
                    '* <a href="https://towardsdatascience.com/\
                    feature-engineering-for-machine-learning-3a5e293a5114">\
                    Feature engineer columns</a>\n',
                    '* <a href="https://pandas.pydata.org/\
                    pandas-docs/stable/reference/api/pandas\
                    .DataFrame.fillna.html">\
                    Fill null values</a>\n',
                    "* Remove rows with null or inconsistent data",
                ],
            },
            {
                "cell_type": "code",
                "execution_count": 0,
                "metadata": {},
                "outputs": [],
                "source": [
                    "# Add code here to \
                feature engineer your dataframe"
                ],
            },
            # Train test split
            {
                "cell_type": "markdown",
                "metadata": {},
                "source": [
                    "## Train Test Split\n",
                    "\n",
                    "We need to split the data into\
                     validation and train data.\
                     Test data may never touch the model until\
                      we are validating it. We need to break the \
                      x set, which is the data we will use to support\
                       the estimator, and the y data, which is the target\
                        value required to estimate.",
                ],
            },
            {
                "cell_type": "code",
                "execution_count": 0,
                "metadata": {},
                "outputs": [],
                "source": [
                    "from sklearn.model_selection \
                    import train_test_split\n",
                    "\n",
                    "# Adjust test size accordingly\n",
                    "test_size = 0.3\n",
                    "\n",
                    "x = data.drop(['<YOUR TARGET COLUMN>'], axis=1)\n",
                    "y = data('<YOUR TARGET COLUMN>')\n",
                    "x_train, x_test, y_train, \
                    y_test = train_test_split(x, y, test_size=test_size)",
                ],
            },
            # Modeling pipeline
            {
                "cell_type": "markdown",
                "metadata": {},
                "source": [
                    "## Normalize the data and train the model\n",
                    "\n",
                    "Dimensions with higher magnitude have a greater\
                     influence on the estimation than those with lower \
                     magnitude. Hence, it is crucial to normalize the\
                      scale of the datapoints columns. \n",
                    '<a href="https://scikit-learn.org/\
                    stable/auto_examples/preprocessing/plot_all_scaling.html">\
                    It is possible to play with other scalers as well</a>\n',
                    "\n",
                    "After scaling, it is time to create the model \
                    and train it. Given the formula options, the \
                    estimator below is an adequate pick given your \
                    Ritchie inputs. Nevertheless, you might want to \
                    experiment with other parameters or even estimators as \
                    well <a href='https://scikit-\
                    learn.org/stable/modules/classes.html'>\
                    You should check the \
                    documentation for other estimators</a>\n",
                    "\n",
                    "How about we streamline this process on a pipeline?\
                     Putting it altogether, we \
                     can write the following code:",
                ],
            },
            {
                "cell_type": "code",
                "execution_count": 0,
                "metadata": {},
                "outputs": [],
                "source": [
                    content[estimator]["import"],
                    "\n",
                    "from sklearn.preprocessing \
                    import StandardScaler\n",
                    "from sklearn.pipeline import Pipeline\n",
                    "\n",
                    "scaler = StandardScaler()\n",
                    "model = {}()\n".format(estimator),
                    "\n",
                    "pipe = Pipeline([('scaler', \
                    scaler), ('model', model)])\n",
                    "pipe.fit(x_train, y_train)\n",
                    "## Uncomment to check the model's score\n",
                    "# print(pipe.score(x_test, y_test))",
                ],
            },
            # Metrics section
            *metrics[content[estimator]["evaluation"]],
            # Optional grid search
            *grid_search(estimator),
            # Model persist
            {
                "cell_type": "markdown",
                "metadata": {},
                "source": [
                    "## Persist the model\n",
                    "\n",
                    "Now that the model is finished we can use \
                    `sklearn`'s persistence library `joblib`\
                     which is more efficient than python's default `pickle`",
                ],
            },
            {
                "cell_type": "code",
                "execution_count": 0,
                "metadata": {},
                "outputs": [],
                "source": [
                    "from joblib import dump, load\n",
                    "\n",
                    "# Save the model\n",
                    "dump(pipe, 'filename.joblib') \n",
                    "\n",
                    "# Load the model\n",
                    "model = load('filename.joblib') ",
                ],
            },
        ],
        "metadata": {
            "kernelspec": {
                "display_name": "Python 3",
                "language": "python",
                "name": "python3",
            },
            "language_info": {
                "codemirror_mode": {"name": "ipython", "version": 3},
                "file_extension": ".py",
                "mimetype": "text/x-python",
                "name": "python",
                "nbconvert_exporter": "python",
                "pygments_lexer": "ipython3",
                "version": "3.7.6",
            },
        },
        "nbformat": 4,
        "nbformat_minor": 1,
    }
