# NoDubs

A really dumb but useful HTTP redirector service. 

`http[s]://www.$example.com => https://$example.com` for anyone who points a 
`www.`-prefixed domain at the server's IP. 

## Status

This is an experiment. You absolutely shouldn't use it for anything serious.

## Motivation

People have been trained to type "www.example.com", and I'm always having to add
some mechanism to redirect requests to "example.com". That used to be easy when
you had a traditional web server (apache/nginx/caddy), but has gotten annoying
in the PaaS era.

## Operation

This stupid bit of software does that for you with zero configuration. You just
point www.yourdomain.com at a server running nodubs, and "Bob's your uncle", as 
the kids almost certainly never say any more.


