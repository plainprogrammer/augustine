# Augustine

Augustine is an experiment to build a very basic MVCish application
framework in Go.

## What Works?

**NOT MUCH!** Right now there is a really primitive notion of
middlewares and a middleware stack. Inspiration came from Ruby's Rack
framework.

## What's Missing?

**A LOT!** Augustine has not higher level notions of controllers, no
concept of models and no distinct view layer. It also lacks a proper
router for the moment, although that's next on the list of experiments
to implement.
