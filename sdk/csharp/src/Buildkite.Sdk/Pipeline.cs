using System.Collections.Generic;
using System.Text.Json;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;

namespace Buildkite.Sdk
{
    public class Pipeline
    {
        public Dictionary<string, object>? Agents { get; set; }
        public Dictionary<string, object>? Env { get; set; }
        public List<object>? Notify { get; set; }
        public List<PipelineStep> Steps { get; set; } = new List<PipelineStep>();

        /// <summary>
        /// Add an agent to target by tag
        /// </summary>
        /// <param name="tagName">The agent tag name</param>
        /// <param name="tagValue">The agent tag value</param>
        public void AddAgent(string tagName, object tagValue)
        {
            Agents ??= new Dictionary<string, object>();
            Agents[tagName] = tagValue;
        }

        /// <summary>
        /// Add an environment variable
        /// </summary>
        /// <param name="key">The environment variable key</param>
        /// <param name="value">The environment variable value</param>
        public void AddEnvironmentVariable(string key, object value)
        {
            Env ??= new Dictionary<string, object>();
            Env[key] = value;
        }

        /// <summary>
        /// Add a notification
        /// </summary>
        /// <param name="notify">The notification configuration</param>
        public void AddNotify(object notify)
        {
            Notify ??= new List<object>();
            Notify.Add(notify);
        }

        /// <summary>
        /// Add a step to the pipeline
        /// </summary>
        /// <param name="step">The pipeline step to add</param>
        /// <returns>The pipeline instance for method chaining</returns>
        public Pipeline AddStep(PipelineStep step)
        {
            Steps.Add(step);
            return this;
        }

        private Dictionary<string, object> Build()
        {
            var pipeline = new Dictionary<string, object>();

            if (Steps.Count > 0)
            {
                pipeline["steps"] = Steps;
            }

            if (Agents != null && Agents.Count > 0)
            {
                pipeline["agents"] = Agents;
            }

            if (Env != null && Env.Count > 0)
            {
                pipeline["env"] = Env;
            }

            if (Notify != null && Notify.Count > 0)
            {
                pipeline["notify"] = Notify;
            }

            return pipeline;
        }

        /// <summary>
        /// Serialize the pipeline as a JSON string
        /// </summary>
        /// <returns>JSON representation of the pipeline</returns>
        public string ToJson()
        {
            var settings = new JsonSerializerSettings
            {
                Formatting = Formatting.Indented,
                ContractResolver = new CamelCasePropertyNamesContractResolver(),
                NullValueHandling = NullValueHandling.Ignore
            };
            return JsonConvert.SerializeObject(Build(), settings);
        }

        /// <summary>
        /// Serialize the pipeline as a YAML string
        /// </summary>
        /// <returns>YAML representation of the pipeline</returns>
        public string ToYaml()
        {
            var serializer = new SerializerBuilder()
                .WithNamingConvention(CamelCaseNamingConvention.Instance)
                .Build();
            return serializer.Serialize(Build());
        }
    }
}