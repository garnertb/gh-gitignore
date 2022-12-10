# gh gitignore

This `gh` extension enables developers to quickly bootstrap new projects with the appropriate `.gitignore` through a CLI interface.

## Examples

Bootstrap a node project with the [nodejs gitignore](https://github.com/github/gitignore/blob/main/Node.gitignore).

```bash
gh gitignore node > .gitignore
```

## What types of projects does this extension support?

```bash
Usage: gh-gitignore [options] [command]

Options:
  -h, --help             display help for command

Commands:
  al                     get the al .gitignore file
  actionscript           get the actionscript .gitignore file
  ada                    get the ada .gitignore file
  agda                   get the agda .gitignore file
  android                get the android .gitignore file
  appengine              get the appengine .gitignore file
  appceleratortitanium   get the appceleratortitanium .gitignore file
  archlinuxpackages      get the archlinuxpackages .gitignore file
  autotools              get the autotools .gitignore file
  c++                    get the c++ .gitignore file
  c                      get the c .gitignore file
  cfwheels               get the cfwheels .gitignore file
  cmake                  get the cmake .gitignore file
  cuda                   get the cuda .gitignore file
  cakephp                get the cakephp .gitignore file
  chefcookbook           get the chefcookbook .gitignore file
  clojure                get the clojure .gitignore file
  codeigniter            get the codeigniter .gitignore file
  commonlisp             get the commonlisp .gitignore file
  composer               get the composer .gitignore file
  concrete5              get the concrete5 .gitignore file
  coq                    get the coq .gitignore file
  craftcms               get the craftcms .gitignore file
  d                      get the d .gitignore file
  dm                     get the dm .gitignore file
  dart                   get the dart .gitignore file
  delphi                 get the delphi .gitignore file
  drupal                 get the drupal .gitignore file
  episerver              get the episerver .gitignore file
  eagle                  get the eagle .gitignore file
  elisp                  get the elisp .gitignore file
  elixir                 get the elixir .gitignore file
  elm                    get the elm .gitignore file
  erlang                 get the erlang .gitignore file
  expressionengine       get the expressionengine .gitignore file
  extjs                  get the extjs .gitignore file
  fancy                  get the fancy .gitignore file
  finale                 get the finale .gitignore file
  flaxengine             get the flaxengine .gitignore file
  forcedotcom            get the forcedotcom .gitignore file
  fortran                get the fortran .gitignore file
  fuelphp                get the fuelphp .gitignore file
  gwt                    get the gwt .gitignore file
  gcov                   get the gcov .gitignore file
  gitbook                get the gitbook .gitignore file
  go                     get the go .gitignore file
  godot                  get the godot .gitignore file
  gradle                 get the gradle .gitignore file
  grails                 get the grails .gitignore file
  haskell                get the haskell .gitignore file
  igorpro                get the igorpro .gitignore file
  idris                  get the idris .gitignore file
  jboss                  get the jboss .gitignore file
  jenkins_home           get the jenkins_home .gitignore file
  java                   get the java .gitignore file
  jekyll                 get the jekyll .gitignore file
  joomla                 get the joomla .gitignore file
  julia                  get the julia .gitignore file
  kicad                  get the kicad .gitignore file
  kohana                 get the kohana .gitignore file
  kotlin                 get the kotlin .gitignore file
  labview                get the labview .gitignore file
  laravel                get the laravel .gitignore file
  leiningen              get the leiningen .gitignore file
  lemonstand             get the lemonstand .gitignore file
  lilypond               get the lilypond .gitignore file
  lithium                get the lithium .gitignore file
  lua                    get the lua .gitignore file
  magento                get the magento .gitignore file
  maven                  get the maven .gitignore file
  mercury                get the mercury .gitignore file
  metaprogrammingsystem  get the metaprogrammingsystem .gitignore file
  nanoc                  get the nanoc .gitignore file
  nim                    get the nim .gitignore file
  node                   get the node .gitignore file
  ocaml                  get the ocaml .gitignore file
  objective-c            get the objective-c .gitignore file
  opa                    get the opa .gitignore file
  opencart               get the opencart .gitignore file
  oracleforms            get the oracleforms .gitignore file
  packer                 get the packer .gitignore file
  perl                   get the perl .gitignore file
  phalcon                get the phalcon .gitignore file
  playframework          get the playframework .gitignore file
  plone                  get the plone .gitignore file
  prestashop             get the prestashop .gitignore file
  processing             get the processing .gitignore file
  purescript             get the purescript .gitignore file
  python                 get the python .gitignore file
  qooxdoo                get the qooxdoo .gitignore file
  qt                     get the qt .gitignore file
  r                      get the r .gitignore file
  ros                    get the ros .gitignore file
  racket                 get the racket .gitignore file
  rails                  get the rails .gitignore file
  raku                   get the raku .gitignore file
  rhodesrhomobile        get the rhodesrhomobile .gitignore file
  ruby                   get the ruby .gitignore file
  rust                   get the rust .gitignore file
  scons                  get the scons .gitignore file
  sass                   get the sass .gitignore file
  scala                  get the scala .gitignore file
  scheme                 get the scheme .gitignore file
  scrivener              get the scrivener .gitignore file
  sdcc                   get the sdcc .gitignore file
  seamgen                get the seamgen .gitignore file
  sketchup               get the sketchup .gitignore file
  smalltalk              get the smalltalk .gitignore file
  stella                 get the stella .gitignore file
  sugarcrm               get the sugarcrm .gitignore file
  swift                  get the swift .gitignore file
  symfony                get the symfony .gitignore file
  symphonycms            get the symphonycms .gitignore file
  tex                    get the tex .gitignore file
  terraform              get the terraform .gitignore file
  textpattern            get the textpattern .gitignore file
  turbogears2            get the turbogears2 .gitignore file
  twincat3               get the twincat3 .gitignore file
  typo3                  get the typo3 .gitignore file
  unity                  get the unity .gitignore file
  unrealengine           get the unrealengine .gitignore file
  vvvv                   get the vvvv .gitignore file
  visualstudio           get the visualstudio .gitignore file
  waf                    get the waf .gitignore file
  wordpress              get the wordpress .gitignore file
  xojo                   get the xojo .gitignore file
  yeoman                 get the yeoman .gitignore file
  yii                    get the yii .gitignore file
  zendframework          get the zendframework .gitignore file
  zephir                 get the zephir .gitignore file
  help [command]         display help for command
```