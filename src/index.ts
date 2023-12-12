import { Construct } from 'constructs';
import * as k8s from './imports/k8s';

export class Hello {
  public sayHello() {
    return 'hello, world!';
  }
}

export interface CustomizedConfigmapOptions {
  readonly namespace?: string;
  readonly data: { [key: string]: string };
}

export class CustomizedConfigmap extends Construct {
  public readonly name?: string;
  public readonly namespace?: string;

  constructor(scope: Construct, name: string, opts: CustomizedConfigmapOptions) {
    super(scope, name);
    const namespace = opts.namespace ?? 'default';
    this.namespace = namespace;
    const data = opts.data;

    const configmapOpts: k8s.KubeConfigMapProps = {
      metadata: {
        name: this.namespace,
        namespace: this.namespace,
      },
      data: data,
    };
    const cfm = new k8s.KubeConfigMap(scope, 'customized-configmap', configmapOpts);
    this.name = cfm.name;
  }
}

export interface ResourceQuantity {
  readonly cpu?: string;
  readonly memory?: string;
}

export function ConvertQuantity(user: ResourceQuantity | undefined, defaults: { cpu: string; memory: string }): { [key: string]: k8s.Quantity } {
  // defaults
  if (!user) {
    return {
      cpu: k8s.Quantity.fromString(defaults.cpu),
      memory: k8s.Quantity.fromString(defaults.memory),
    };
  }

  const result: { [key: string]: k8s.Quantity } = { };
  if (user.cpu) {
    result.cpu = k8s.Quantity.fromString(user.cpu);
  }
  if (user.memory) {
    result.memory = k8s.Quantity.fromString(user.memory);
  }
  return result;
};

export interface CustomizedDeploymentOptions {
  readonly namespace?: string;
  readonly labels: { [key: string]: string };
  readonly annotations: { [key: string]: string };
  readonly image: string;
  readonly replicas?: number;
  readonly serviceAccountName?: string;
  readonly resources?: ResourceRequirements;
}

export interface ResourceQuantity {
  readonly cpu?: string;
  readonly memory?: string;
}

export interface ResourceRequirements {
  readonly limits?: ResourceQuantity;
  readonly requests?: ResourceQuantity;
}

export class CustomizedDeployment extends Construct {
  public readonly name: string;
  public readonly namespace: string;

  constructor(scope: Construct, name: string, selector: {[key: string]: string}, opts: CustomizedDeploymentOptions) {
    super(scope, name);
    const namespace = opts.namespace ?? 'default';
    this.namespace = namespace;
    this.name = namespace;

    const serviceAccountName = opts.serviceAccountName;
    const replicas = opts.replicas ?? 1;
    const image = opts.image;
    const annotations = opts.annotations;
    const labels = opts.labels;
    const resources = {
      limits: ConvertQuantity(opts.resources?.limits, { cpu: '500m', memory: '250Mi' }),
      requests: ConvertQuantity(opts.resources?.requests, { cpu: '500m', memory: '250Mi' }),
    };

    const deploymentOpts: k8s.KubeDeploymentProps = {
      metadata: {
        name: this.namespace,
        namespace: this.namespace,
        annotations: annotations,
        labels: labels,
      },
      spec: {
        replicas: replicas,
        selector: { matchLabels: selector },
        template: {
          metadata: { labels: selector },
          spec: {
            serviceAccountName: serviceAccountName,
            containers: [{
              name: 'test-app',
              image: image,
              envFrom: [{ configMapRef: { name: 'customized-configmap', optional: true } }],
              resources: resources,
            }],
          },
        },
      },
    };
    new k8s.KubeDeployment(this, 'customized-deployment', deploymentOpts);
  }
}

export interface CustomizedOptions {
  readonly deploymentOptions: CustomizedDeploymentOptions;
  readonly configmapOptions: CustomizedConfigmapOptions;
}

export class CustomizedApp extends Construct {
  constructor(scope: Construct, name: string, opts: CustomizedOptions) {
    super(scope, name);

    const selector = { app: 'test-app' };
    new CustomizedDeployment(this, 'deployment', selector, opts.deploymentOptions);
    new CustomizedConfigmap(this, 'configmap', opts.configmapOptions);
  }
}