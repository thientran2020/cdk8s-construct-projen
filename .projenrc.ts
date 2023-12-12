import { cdk8s } from 'projen';
const project = new cdk8s.ConstructLibraryCdk8s({
  author: 'thientran',
  authorAddress: 'thienytran14@gmail.com',
  cdk8sVersion: '2.68.18',
  constructsVersion: '10.0.0',
  defaultReleaseBranch: 'main',
  jsiiVersion: '~5.0.0',
  name: 'cdk8s-custom-construct',
  projenrcTs: true,
  repositoryUrl: 'https://github.com/thientran2020/cdk8s-construct-projen.git',
  publishToGo: {
    moduleName: 'github.com/thientran2020/cdk8s-construct-projen',
    githubTokenSecret: 'GO_PUBLISH_TOKEN',
  },

  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();