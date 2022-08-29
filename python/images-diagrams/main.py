from diagrams import Cluster, Diagram
from diagrams.programming.framework import React
from diagrams.programming.language import Swift, Kotlin
from diagrams.aws.network import CF
from diagrams.generic.blank import Blank

with Diagram('Architecture of images'):

    with Cluster('Application'):
        app2 = Swift('')
        app = React('')
        app3 = Kotlin('')
    with Cluster('System'):
        app = React('')


    cf = CF('CloudFront')
    app >> cf
    with: Cluster('Application') >> cf
