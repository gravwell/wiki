# if this fails to build, you may need to install the pandoc xnos suite:
# pip install pandoc-fignos pandoc-eqnos pandoc-tablenos pandoc-secnos --user

pdf: migrate.pandoc
	pandoc migrate.pandoc --number-sections --filter pandoc-xnos -s -o migrate.pdf

wiki: migrate.pandoc
	echo '# Migration Tool' > migrate.md
	pandoc -f markdown migrate.pandoc --wrap=none -t gfm -s --shift-heading-level=1 >> migrate.md