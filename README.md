# ESV API for Go

This is a small SDK for allowing Go programs to easily contact the ESV Bible API
at <https://api.esv.org/>.

# Implemented vs. To Do

The passage text, passage HTML, and search endpoints are implemented.

The audio endpoint is not working. However, that it should be very simple to
provide the audio link given by the redirect. I just haven't written the code
for that into the generator template.
