# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CustomizedApp <a name="CustomizedApp" id="cdk8s-custom-construct.CustomizedApp"></a>

#### Initializers <a name="Initializers" id="cdk8s-custom-construct.CustomizedApp.Initializer"></a>

```typescript
import { CustomizedApp } from 'cdk8s-custom-construct'

new CustomizedApp(scope: Construct, name: string, opts: CustomizedOptions)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedApp.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedApp.Initializer.parameter.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedApp.Initializer.parameter.opts">opts</a></code> | <code><a href="#cdk8s-custom-construct.CustomizedOptions">CustomizedOptions</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk8s-custom-construct.CustomizedApp.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `name`<sup>Required</sup> <a name="name" id="cdk8s-custom-construct.CustomizedApp.Initializer.parameter.name"></a>

- *Type:* string

---

##### `opts`<sup>Required</sup> <a name="opts" id="cdk8s-custom-construct.CustomizedApp.Initializer.parameter.opts"></a>

- *Type:* <a href="#cdk8s-custom-construct.CustomizedOptions">CustomizedOptions</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedApp.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk8s-custom-construct.CustomizedApp.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.




### CustomizedConfigmap <a name="CustomizedConfigmap" id="cdk8s-custom-construct.CustomizedConfigmap"></a>

#### Initializers <a name="Initializers" id="cdk8s-custom-construct.CustomizedConfigmap.Initializer"></a>

```typescript
import { CustomizedConfigmap } from 'cdk8s-custom-construct'

new CustomizedConfigmap(scope: Construct, name: string, opts: CustomizedConfigmapOptions)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedConfigmap.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedConfigmap.Initializer.parameter.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedConfigmap.Initializer.parameter.opts">opts</a></code> | <code><a href="#cdk8s-custom-construct.CustomizedConfigmapOptions">CustomizedConfigmapOptions</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk8s-custom-construct.CustomizedConfigmap.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `name`<sup>Required</sup> <a name="name" id="cdk8s-custom-construct.CustomizedConfigmap.Initializer.parameter.name"></a>

- *Type:* string

---

##### `opts`<sup>Required</sup> <a name="opts" id="cdk8s-custom-construct.CustomizedConfigmap.Initializer.parameter.opts"></a>

- *Type:* <a href="#cdk8s-custom-construct.CustomizedConfigmapOptions">CustomizedConfigmapOptions</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedConfigmap.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk8s-custom-construct.CustomizedConfigmap.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.


#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedConfigmap.property.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedConfigmap.property.namespace">namespace</a></code> | <code>string</code> | *No description.* |

---

##### `name`<sup>Optional</sup> <a name="name" id="cdk8s-custom-construct.CustomizedConfigmap.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

---

##### `namespace`<sup>Optional</sup> <a name="namespace" id="cdk8s-custom-construct.CustomizedConfigmap.property.namespace"></a>

```typescript
public readonly namespace: string;
```

- *Type:* string

---


### CustomizedDeployment <a name="CustomizedDeployment" id="cdk8s-custom-construct.CustomizedDeployment"></a>

#### Initializers <a name="Initializers" id="cdk8s-custom-construct.CustomizedDeployment.Initializer"></a>

```typescript
import { CustomizedDeployment } from 'cdk8s-custom-construct'

new CustomizedDeployment(scope: Construct, name: string, selector: {[ key: string ]: string}, opts: CustomizedDeploymentOptions)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedDeployment.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeployment.Initializer.parameter.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeployment.Initializer.parameter.selector">selector</a></code> | <code>{[ key: string ]: string}</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeployment.Initializer.parameter.opts">opts</a></code> | <code><a href="#cdk8s-custom-construct.CustomizedDeploymentOptions">CustomizedDeploymentOptions</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk8s-custom-construct.CustomizedDeployment.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `name`<sup>Required</sup> <a name="name" id="cdk8s-custom-construct.CustomizedDeployment.Initializer.parameter.name"></a>

- *Type:* string

---

##### `selector`<sup>Required</sup> <a name="selector" id="cdk8s-custom-construct.CustomizedDeployment.Initializer.parameter.selector"></a>

- *Type:* {[ key: string ]: string}

---

##### `opts`<sup>Required</sup> <a name="opts" id="cdk8s-custom-construct.CustomizedDeployment.Initializer.parameter.opts"></a>

- *Type:* <a href="#cdk8s-custom-construct.CustomizedDeploymentOptions">CustomizedDeploymentOptions</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedDeployment.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk8s-custom-construct.CustomizedDeployment.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.


#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedDeployment.property.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeployment.property.namespace">namespace</a></code> | <code>string</code> | *No description.* |

---

##### `name`<sup>Required</sup> <a name="name" id="cdk8s-custom-construct.CustomizedDeployment.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

---

##### `namespace`<sup>Required</sup> <a name="namespace" id="cdk8s-custom-construct.CustomizedDeployment.property.namespace"></a>

```typescript
public readonly namespace: string;
```

- *Type:* string

---


## Structs <a name="Structs" id="Structs"></a>

### CustomizedConfigmapOptions <a name="CustomizedConfigmapOptions" id="cdk8s-custom-construct.CustomizedConfigmapOptions"></a>

#### Initializer <a name="Initializer" id="cdk8s-custom-construct.CustomizedConfigmapOptions.Initializer"></a>

```typescript
import { CustomizedConfigmapOptions } from 'cdk8s-custom-construct'

const customizedConfigmapOptions: CustomizedConfigmapOptions = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedConfigmapOptions.property.data">data</a></code> | <code>{[ key: string ]: string}</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedConfigmapOptions.property.namespace">namespace</a></code> | <code>string</code> | *No description.* |

---

##### `data`<sup>Required</sup> <a name="data" id="cdk8s-custom-construct.CustomizedConfigmapOptions.property.data"></a>

```typescript
public readonly data: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

---

##### `namespace`<sup>Optional</sup> <a name="namespace" id="cdk8s-custom-construct.CustomizedConfigmapOptions.property.namespace"></a>

```typescript
public readonly namespace: string;
```

- *Type:* string

---

### CustomizedDeploymentOptions <a name="CustomizedDeploymentOptions" id="cdk8s-custom-construct.CustomizedDeploymentOptions"></a>

#### Initializer <a name="Initializer" id="cdk8s-custom-construct.CustomizedDeploymentOptions.Initializer"></a>

```typescript
import { CustomizedDeploymentOptions } from 'cdk8s-custom-construct'

const customizedDeploymentOptions: CustomizedDeploymentOptions = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedDeploymentOptions.property.annotations">annotations</a></code> | <code>{[ key: string ]: string}</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeploymentOptions.property.image">image</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeploymentOptions.property.labels">labels</a></code> | <code>{[ key: string ]: string}</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeploymentOptions.property.namespace">namespace</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeploymentOptions.property.replicas">replicas</a></code> | <code>number</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeploymentOptions.property.resources">resources</a></code> | <code><a href="#cdk8s-custom-construct.ResourceRequirements">ResourceRequirements</a></code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedDeploymentOptions.property.serviceAccountName">serviceAccountName</a></code> | <code>string</code> | *No description.* |

---

##### `annotations`<sup>Required</sup> <a name="annotations" id="cdk8s-custom-construct.CustomizedDeploymentOptions.property.annotations"></a>

```typescript
public readonly annotations: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

---

##### `image`<sup>Required</sup> <a name="image" id="cdk8s-custom-construct.CustomizedDeploymentOptions.property.image"></a>

```typescript
public readonly image: string;
```

- *Type:* string

---

##### `labels`<sup>Required</sup> <a name="labels" id="cdk8s-custom-construct.CustomizedDeploymentOptions.property.labels"></a>

```typescript
public readonly labels: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}

---

##### `namespace`<sup>Optional</sup> <a name="namespace" id="cdk8s-custom-construct.CustomizedDeploymentOptions.property.namespace"></a>

```typescript
public readonly namespace: string;
```

- *Type:* string

---

##### `replicas`<sup>Optional</sup> <a name="replicas" id="cdk8s-custom-construct.CustomizedDeploymentOptions.property.replicas"></a>

```typescript
public readonly replicas: number;
```

- *Type:* number

---

##### `resources`<sup>Optional</sup> <a name="resources" id="cdk8s-custom-construct.CustomizedDeploymentOptions.property.resources"></a>

```typescript
public readonly resources: ResourceRequirements;
```

- *Type:* <a href="#cdk8s-custom-construct.ResourceRequirements">ResourceRequirements</a>

---

##### `serviceAccountName`<sup>Optional</sup> <a name="serviceAccountName" id="cdk8s-custom-construct.CustomizedDeploymentOptions.property.serviceAccountName"></a>

```typescript
public readonly serviceAccountName: string;
```

- *Type:* string

---

### CustomizedOptions <a name="CustomizedOptions" id="cdk8s-custom-construct.CustomizedOptions"></a>

#### Initializer <a name="Initializer" id="cdk8s-custom-construct.CustomizedOptions.Initializer"></a>

```typescript
import { CustomizedOptions } from 'cdk8s-custom-construct'

const customizedOptions: CustomizedOptions = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.CustomizedOptions.property.configmapOptions">configmapOptions</a></code> | <code><a href="#cdk8s-custom-construct.CustomizedConfigmapOptions">CustomizedConfigmapOptions</a></code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.CustomizedOptions.property.deploymentOptions">deploymentOptions</a></code> | <code><a href="#cdk8s-custom-construct.CustomizedDeploymentOptions">CustomizedDeploymentOptions</a></code> | *No description.* |

---

##### `configmapOptions`<sup>Required</sup> <a name="configmapOptions" id="cdk8s-custom-construct.CustomizedOptions.property.configmapOptions"></a>

```typescript
public readonly configmapOptions: CustomizedConfigmapOptions;
```

- *Type:* <a href="#cdk8s-custom-construct.CustomizedConfigmapOptions">CustomizedConfigmapOptions</a>

---

##### `deploymentOptions`<sup>Required</sup> <a name="deploymentOptions" id="cdk8s-custom-construct.CustomizedOptions.property.deploymentOptions"></a>

```typescript
public readonly deploymentOptions: CustomizedDeploymentOptions;
```

- *Type:* <a href="#cdk8s-custom-construct.CustomizedDeploymentOptions">CustomizedDeploymentOptions</a>

---

### ResourceQuantity <a name="ResourceQuantity" id="cdk8s-custom-construct.ResourceQuantity"></a>

#### Initializer <a name="Initializer" id="cdk8s-custom-construct.ResourceQuantity.Initializer"></a>

```typescript
import { ResourceQuantity } from 'cdk8s-custom-construct'

const resourceQuantity: ResourceQuantity = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.ResourceQuantity.property.cpu">cpu</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.ResourceQuantity.property.memory">memory</a></code> | <code>string</code> | *No description.* |

---

##### `cpu`<sup>Optional</sup> <a name="cpu" id="cdk8s-custom-construct.ResourceQuantity.property.cpu"></a>

```typescript
public readonly cpu: string;
```

- *Type:* string

---

##### `memory`<sup>Optional</sup> <a name="memory" id="cdk8s-custom-construct.ResourceQuantity.property.memory"></a>

```typescript
public readonly memory: string;
```

- *Type:* string

---

### ResourceRequirements <a name="ResourceRequirements" id="cdk8s-custom-construct.ResourceRequirements"></a>

#### Initializer <a name="Initializer" id="cdk8s-custom-construct.ResourceRequirements.Initializer"></a>

```typescript
import { ResourceRequirements } from 'cdk8s-custom-construct'

const resourceRequirements: ResourceRequirements = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk8s-custom-construct.ResourceRequirements.property.limits">limits</a></code> | <code><a href="#cdk8s-custom-construct.ResourceQuantity">ResourceQuantity</a></code> | *No description.* |
| <code><a href="#cdk8s-custom-construct.ResourceRequirements.property.requests">requests</a></code> | <code><a href="#cdk8s-custom-construct.ResourceQuantity">ResourceQuantity</a></code> | *No description.* |

---

##### `limits`<sup>Optional</sup> <a name="limits" id="cdk8s-custom-construct.ResourceRequirements.property.limits"></a>

```typescript
public readonly limits: ResourceQuantity;
```

- *Type:* <a href="#cdk8s-custom-construct.ResourceQuantity">ResourceQuantity</a>

---

##### `requests`<sup>Optional</sup> <a name="requests" id="cdk8s-custom-construct.ResourceRequirements.property.requests"></a>

```typescript
public readonly requests: ResourceQuantity;
```

- *Type:* <a href="#cdk8s-custom-construct.ResourceQuantity">ResourceQuantity</a>

---

## Classes <a name="Classes" id="Classes"></a>

### Hello <a name="Hello" id="cdk8s-custom-construct.Hello"></a>

#### Initializers <a name="Initializers" id="cdk8s-custom-construct.Hello.Initializer"></a>

```typescript
import { Hello } from 'cdk8s-custom-construct'

new Hello()
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk8s-custom-construct.Hello.sayHello">sayHello</a></code> | *No description.* |

---

##### `sayHello` <a name="sayHello" id="cdk8s-custom-construct.Hello.sayHello"></a>

```typescript
public sayHello(): string
```





