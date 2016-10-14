### PLAYING WITH THE EXPERT SYSTEMS API

#### ENDPOINTS:

http://cogito-api.cloudapp.net:8091/cogito/v1/annotation/annotate/economist-plan.xml

This endpoint produces the full semantic analysis results: categorization (Media Topic taxonomy), extraction (relevant terms and concepts), default mining (entities and tags), main sentences and topics.

http://cogito-api.cloudapp.net:8091/cogito/v1/annotation/annotate/economist-mediatopic.xml
        
This endpoint produces the categorization (Media Topic taxonomy) results only.
     
http://cogito-api.cloudapp.net:8091/cogito/v1/annotation/annotate/economist-cogito-synthesis.xml
       
This endpoint produces results from extraction (relevant terms and concepts), default mining (entities and tags), main sentences and topics.

#### RUNNING THE ANALYSIS

* Fill up the samples folder with content from our nodes (nid.sample).
* Build the binary.
* Run ```./expertapi --nid=<nid> --mode=<mode>```.
* Mode defaults to full analysis, specify ```categorization``` or ```extraction_default_mining``` to hit the other endpoints.
* The system will write the output into the output folder prefixed by result type and timestamp.

