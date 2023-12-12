import { cdk8s } from 'projen';
const project = new cdk8s.ConstructLibraryCdk8s({
  author: 'thientran',
  authorAddress: 'thienytran14@gmail.com',
  cdk8sVersion: '1.4.10',
  defaultReleaseBranch: 'main',
  jsiiVersion: '~5.0.0',
  name: 'cdk8s-custom-construct',
  projenrcTs: true,
  repositoryUrl: 'https://github.com/thienytran14/cdk8s-custom-construct.git',

  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();