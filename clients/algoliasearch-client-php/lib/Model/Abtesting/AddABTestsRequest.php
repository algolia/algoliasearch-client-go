<?php

// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

namespace Algolia\AlgoliaSearch\Model\Abtesting;

/**
 * AddABTestsRequest Class Doc Comment
 *
 * @category Class
 * @package Algolia\AlgoliaSearch
 */
class AddABTestsRequest extends \Algolia\AlgoliaSearch\Model\AbstractModel implements
    ModelInterface,
    \ArrayAccess,
    \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'name' => 'string',
        'variant' => '\Algolia\AlgoliaSearch\Model\Abtesting\AddABTestsVariant[]',
        'endAt' => 'string',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'name' => null,
        'variant' => null,
        'endAt' => null,
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelTypes()
    {
        return self::$modelTypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelFormats()
    {
        return self::$modelFormats;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'name' => 'setName',
        'variant' => 'setVariant',
        'endAt' => 'setEndAt',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'name' => 'getName',
        'variant' => 'getVariant',
        'endAt' => 'getEndAt',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     */
    public function __construct(array $data = null)
    {
        if (isset($data['name'])) {
            $this->container['name'] = $data['name'];
        }
        if (isset($data['variant'])) {
            $this->container['variant'] = $data['variant'];
        }
        if (isset($data['endAt'])) {
            $this->container['endAt'] = $data['endAt'];
        }
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if (
            !isset($this->container['name']) ||
            $this->container['name'] === null
        ) {
            $invalidProperties[] = "'name' can't be null";
        }
        if (
            !isset($this->container['variant']) ||
            $this->container['variant'] === null
        ) {
            $invalidProperties[] = "'variant' can't be null";
        }
        if (count($this->container['variant']) > 2) {
            $invalidProperties[] =
                "invalid value for 'variant', number of items must be less than or equal to 2.";
        }

        if (count($this->container['variant']) < 2) {
            $invalidProperties[] =
                "invalid value for 'variant', number of items must be greater than or equal to 2.";
        }

        if (
            !isset($this->container['endAt']) ||
            $this->container['endAt'] === null
        ) {
            $invalidProperties[] = "'endAt' can't be null";
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }

    /**
     * Gets name
     *
     * @return string
     */
    public function getName()
    {
        return $this->container['name'] ?? null;
    }

    /**
     * Sets name
     *
     * @param string $name A/B test name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets variant
     *
     * @return \Algolia\AlgoliaSearch\Model\Abtesting\AddABTestsVariant[]
     */
    public function getVariant()
    {
        return $this->container['variant'] ?? null;
    }

    /**
     * Sets variant
     *
     * @param \Algolia\AlgoliaSearch\Model\Abtesting\AddABTestsVariant[] $variant list of 2 variants for the A/B test
     *
     * @return self
     */
    public function setVariant($variant)
    {
        if (count($variant) > 2) {
            throw new \InvalidArgumentException(
                'invalid value for $variant when calling AddABTestsRequest., number of items must be less than or equal to 2.'
            );
        }
        if (count($variant) < 2) {
            throw new \InvalidArgumentException(
                'invalid length for $variant when calling AddABTestsRequest., number of items must be greater than or equal to 2.'
            );
        }
        $this->container['variant'] = $variant;

        return $this;
    }

    /**
     * Gets endAt
     *
     * @return string
     */
    public function getEndAt()
    {
        return $this->container['endAt'] ?? null;
    }

    /**
     * Sets endAt
     *
     * @param string $endAt end date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ
     *
     * @return self
     */
    public function setEndAt($endAt)
    {
        $this->container['endAt'] = $endAt;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param int $offset Offset
     *
     * @return bool
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param int $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param int $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }
}