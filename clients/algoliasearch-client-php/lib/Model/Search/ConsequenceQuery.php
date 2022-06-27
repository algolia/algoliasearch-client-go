<?php

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * ConsequenceQuery Class Doc Comment
 *
 * @category Class
 * @package Algolia\AlgoliaSearch
 */
class ConsequenceQuery extends \Algolia\AlgoliaSearch\Model\AbstractModel implements
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
        'remove' => 'string[]',
        'edits' => '\Algolia\AlgoliaSearch\Model\Search\Edit[]',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'remove' => null,
        'edits' => null,
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
        'remove' => 'setRemove',
        'edits' => 'setEdits',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'remove' => 'getRemove',
        'edits' => 'getEdits',
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
        if (isset($data['remove'])) {
            $this->container['remove'] = $data['remove'];
        }
        if (isset($data['edits'])) {
            $this->container['edits'] = $data['edits'];
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
     * Gets remove
     *
     * @return string[]|null
     */
    public function getRemove()
    {
        return $this->container['remove'] ?? null;
    }

    /**
     * Sets remove
     *
     * @param string[]|null $remove words to remove
     *
     * @return self
     */
    public function setRemove($remove)
    {
        $this->container['remove'] = $remove;

        return $this;
    }

    /**
     * Gets edits
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\Edit[]|null
     */
    public function getEdits()
    {
        return $this->container['edits'] ?? null;
    }

    /**
     * Sets edits
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\Edit[]|null $edits edits to apply
     *
     * @return self
     */
    public function setEdits($edits)
    {
        $this->container['edits'] = $edits;

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
