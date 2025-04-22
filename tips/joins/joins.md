# Joins

Joining data is a powerful way to bring together multiple datasets, enrich data based off of sumarizations of other data sets or parse a single dataset multiple times with a single output. While joins are flexible you do need to keep in mind that every joined query will act as a full execution of a query. 



## Example 1 - Enriching data

For this example we want to use the intel data below to filter our web login data to see if any of the logins were from risky IPs. Normally we would have a resource to pull from but this data changes so rapidly that we have ingested it and need to query to find the latest risk.

Below is a basic query of the data but we need to get the latest by IP/

![Intel Data](image.png)

This query will give us the latest by ip and show us the fields we want to work with.

![Latest Risk by IP](image-1.png)

Below is the web login data

![alt text](image-2.png)

Now we can use a compound query to find the latest ip risk from the intel data and join it to the web login data.  If there are any matches they will show data in the `risk` `threat` and `indicator` columns. Check out [compound queries](https://docs.gravwell.io/search/spec.html#compound-queries) to read up on syntax and flags. 


![Matched Intel Data](image-3.png)

This is great but we only want to filter out logins that did not have a match with the threat intel data. For this we can add the `-s` flag to make the lookup strict, which would require the match to be successful to show. 

![Matched Threat Intel](image-4.png)

### Final Query
```
@intel{
	tag=intel ax
	| last ip
	| table ip risk threat indicator
};

tag=web ax
| lookup -s -r @intel srcip ip (risk threat indicator)
| table
```

## Example 2 - Multiple Data Sets

Lets build off the data from the last example. Let's see what happened after a user connected to the risky IP.

If we look at the proxy logs and compare them to the threat intel from our previous example it looks like we have one hit. The user is `drevil` and the internal ip is `192.168.5.25`.

![AV Lookup](image-5.png)

Since that event show that a file was downloaded lets see if there is any activity in the AV tag from that IP. Since we can use compound searches we can keep movething them up the query.

![AV Lookup](image-6.png)

We can see from the results that `drevil` is showing up in a few places lets expand our search to other datasources and match on `IP` or `username`.

![Username Match](image-9.png)

With this result it looks like the user could not get in with the `drevil` credentials. What if they tried more than one username?

![IP Match](image-8.png)

It looks like they use a password spray attack to find an account and performed some SQL table actions. Lets find what they did.

![Results](image-11.png)

The attacker successfully logged in and performed a schema_update with the joenobody account. 


